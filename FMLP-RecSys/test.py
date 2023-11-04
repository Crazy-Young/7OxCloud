import os
import torch
import argparse
import numpy as np

from models import FMLPRecModel
from trainers import FMLPRecTrainer
from utils import EarlyStopping, check_path, set_seed, get_local_time
from utils_demo import get_seq_dic, get_dataloder, get_rating_matrix, generate_rating_matrix_test2, load
from torch.utils.data import DataLoader, RandomSampler, SequentialSampler
from datasets import FMLPRecDataset
import tqdm
import random
import csv


def predict_full(model, seq_out):
    # [item_num hidden_size]
    test_item_emb = model.item_embeddings.weight
    # [batch hidden_size ]
    rating_pred = torch.matmul(seq_out, test_item_emb.transpose(0, 1))
    return rating_pred


def custom_sort(row):
    return (row[0], row[3])


def main(User_idx):
    parser = argparse.ArgumentParser()
    parser.add_argument("--data_dir", default=r"data\\", type=str)
    parser.add_argument("--output_dir", default=r"output\\", type=str)
    parser.add_argument("--data_name", default="video-log", type=str)
    parser.add_argument("--do_eval", action="store_true")
    parser.add_argument("--load_model", default=None, type=str)

    # model args
    parser.add_argument("--model_name", default="FMLPRec", type=str)
    parser.add_argument("--hidden_size", default=16, type=int, help="hidden size of model")
    parser.add_argument("--num_hidden_layers", default=1, type=int, help="number of filter-enhanced blocks")
    parser.add_argument("--num_attention_heads", default=2, type=int)
    parser.add_argument("--hidden_act", default="gelu", type=str)  # gelu relu
    parser.add_argument("--attention_probs_dropout_prob", default=0.5, type=float)
    parser.add_argument("--hidden_dropout_prob", default=0.5, type=float)
    parser.add_argument("--initializer_range", default=0.02, type=float)
    parser.add_argument("--max_seq_length", default=5, type=int)
    parser.add_argument("--no_filters", action="store_true",
                        help="if no filters, filter layers transform to self-attention")

    # train args
    parser.add_argument("--lr", default=0.001, type=float, help="learning rate of adam")
    parser.add_argument("--batch_size", default=16, type=int, help="number of batch_size")
    parser.add_argument("--epochs", default=8, type=int, help="number of epochs")
    parser.add_argument("--no_cuda", action="store_true")
    parser.add_argument("--log_freq", default=1, type=int, help="per epoch print res")
    parser.add_argument("--full_sort", default=True, type=bool)
    parser.add_argument("--patience", default=10, type=int,
                        help="how long to wait after last time validation loss improved")

    parser.add_argument("--seed", default=42, type=int)
    parser.add_argument("--weight_decay", default=0.0, type=float, help="weight_decay of adam")
    parser.add_argument("--adam_beta1", default=0.9, type=float, help="adam first beta value")
    parser.add_argument("--adam_beta2", default=0.999, type=float, help="adam second beta value")
    parser.add_argument("--gpu_id", default="0", type=str, help="gpu_id")
    parser.add_argument("--variance", default=5, type=float)

    args = parser.parse_args()
    args.cuda_condition = torch.cuda.is_available() and not args.no_cuda
    args.device = torch.device("cuda" if args.cuda_condition else "cpu")
    # print(args.device)
    # set_seed(args.seed)
    check_path(args.output_dir)

    os.environ["CUDA_VISIBLE_DEVICES"] = args.gpu_id
    args.cuda_condition = torch.cuda.is_available() and not args.no_cuda

    def find_new_file(dir):
        '''查找目录下最新的文件'''
        file_lists = os.listdir(dir)
        file_lists.sort(key=lambda fn: os.path.getmtime(dir + "\\" + fn)
        if not os.path.isdir(dir + "\\" + fn) else 0)
        # print('最新的文件为： ' + file_lists[-1])
        file = os.path.join(dir, file_lists[-1])
        # print('完整路径：', file)
        return file_lists[-1]

    # 函数调用

    dir = r'output\\'
    str_ = find_new_file(dir)
    ind = str_.find(".")
    args.load_model = str_[:ind]
    # print("load_model:", args.load_model)

    # seq_dic, max_item = get_seq_dic(args)
    args.checkpoint_path = os.path.join(args.output_dir, args.load_model + '.pt')
    if not os.path.exists(args.checkpoint_path):
        print(f"No model input!")
        rec_seq = []
        rec_seq.append([])
        flag = 1
        return rec_seq, flag


    id2index = load("id2index.json")
    args.item_size = len(id2index)  # 需要及时更新
    # print("item_size:", args.item_size)

    # save model args
    cur_time = get_local_time()
    if args.no_filters:
        args.model_name = "SASRec"

    users = []
    users.append(User_idx)
    user_seq = []
    data_file2 = args.data_dir + args.data_name + "-sorted-mapped" + '.json'
    mydict = load(data_file2)
    if (users[0] == 0 or mydict.get(str(users[0])) == None):
        rdint = random.randint(0, len(mydict) - 1)
        # print(len((mydict)))
        users[0] = list(mydict.keys())[rdint]
    # print("user:", users[0])
    user_seq.append(mydict[str(users[0])])
    flag = 0
    # if (len(user_seq[0]) < 1):
    #     flag = 1
    #     rec_seq = []
    #     print("end early")
    #     return rec_seq, flag  # 过小交互记录的用户进行随机推荐

    seq_dic = {'user_seq': user_seq, 'num_users': len(users)}

    test_dataset = FMLPRecDataset(args, seq_dic['user_seq'], data_type='test')
    test_sampler = SequentialSampler(test_dataset)
    test_dataloader = DataLoader(test_dataset, sampler=test_sampler, batch_size=args.batch_size)

    model = FMLPRecModel(args=args)
    if args.cuda_condition:
        model.cuda()

    if args.full_sort:
        test_rating_matrix = generate_rating_matrix_test2(users, seq_dic['user_seq'], seq_dic['num_users'],
                                                          args.item_size)

    if args.load_model is None:
        print(f"No model input!")
        rec_seq = []
        rec_seq.append([])
        flag = 1
        return rec_seq, flag
        exit(0)
    else:

        # trainer.load(args.checkpoint_path)
        original_state_dict = model.state_dict()
        # print(original_state_dict.keys())
        new_dict = torch.load(args.checkpoint_path)
        # print(new_dict.keys())
        for key in new_dict:
            original_state_dict[key] = new_dict[key]
        model.load_state_dict(original_state_dict)
        print(f"Load model from {args.checkpoint_path} for test!")
        # scores, result_info = trainer.test(0, full_sort=args.full_sort)
        model.eval()

        pred_list = None
        dataloader = test_dataloader
        rec_data_iter = tqdm.tqdm(enumerate(dataloader),
                                  desc="Recommendation EP_%s:%d" % ("test", 0),
                                  total=len(dataloader),
                                  bar_format="{l_bar}{r_bar}")

        if args.full_sort:
            # answer_list = None
            args.train_matrix = test_rating_matrix
            for i, batch in rec_data_iter:
                # 0. batch_data will be sent into the device(GPU or cpu)
                batch = tuple(t.to(args.device) for t in batch)
                user_ids, input_ids, answers, _ = batch
                recommend_output = model(input_ids)
                recommend_output = recommend_output[:, -1, :]  # 推荐的结果

                rating_pred = predict_full(model, recommend_output)
                rating_pred = rating_pred.cpu().data.numpy().copy()
                batch_user_index = user_ids.cpu().numpy()
                rating_pred[args.train_matrix[batch_user_index].toarray() > 0] = 0
                # reference: https://stackoverflow.com/a/23734295, https://stackoverflow.com/a/20104162
                # argpartition 时间复杂度O(n)  argsort O(nlogn) 只会做
                # 加负号"-"表示取大的值
                ind = np.argpartition(rating_pred, -20)[:, -20:]
                # 根据返回的下标 从对应维度分别取对应的值 得到每行topk的子表
                arr_ind = rating_pred[np.arange(len(rating_pred))[:, None], ind]
                # 对子表进行排序 得到从大到小的顺序
                arr_ind_argsort = np.argsort(arr_ind)[np.arange(len(rating_pred)), ::-1]
                # 再取一次 从ind中取回 原来的下标
                batch_pred_list = ind[np.arange(len(rating_pred))[:, None], arr_ind_argsort]
                batch_pred_list = batch_pred_list.tolist()
                for i in range(len(batch_pred_list)):
                    demo = list(set(batch_pred_list[i]).difference(set(user_seq[i])))
                    # print("user_Seq:", " ", user_seq)
                    # print("demo:", len(demo), " ", demo)
                    batch_pred_list[i].clear()
                    batch_pred_list[i] = demo

                # print("i:", i, "batch_pred_list:", batch_pred_list)
                index2id = load("index2id.json")
                for i in range(len(batch_pred_list[0])):
                    if index2id.get(str(batch_pred_list[0][i])) == None:
                        continue
                    batch_pred_list[0][i] = int(index2id[str(batch_pred_list[0][i])])
                # print(batch_pred_list)
                return batch_pred_list, flag

# main(User_idx=19)
