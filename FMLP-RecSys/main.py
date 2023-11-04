# @Time   : 2022/2/13
# @Author : Hui Yu
# @Email  : ishyu@outlook.com

import os
import torch
import argparse
import numpy as np

from models import FMLPRecModel
from trainers import FMLPRecTrainer
from utils import EarlyStopping, check_path, set_seed, get_local_time
from utils_demo import get_seq_dic, get_dataloder, get_rating_matrix, load
import csv


def main2():
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

    set_seed(args.seed)
    check_path(args.output_dir)

    os.environ["CUDA_VISIBLE_DEVICES"] = args.gpu_id
    args.cuda_condition = torch.cuda.is_available() and not args.no_cuda
    args.data_file = args.data_dir + args.data_name + '.csv'

    with open(args.data_file, "r", encoding="utf-8") as csvfile:
        csv_reader = csv.reader(csvfile)
        csv_reader = list(csv_reader)
        # print("lenghth:",len(csv_reader))
        if(len(csv_reader)<=2):
            print("data volume too small, out!")
            return 0


    seq_dic, max_item = get_seq_dic(args)

    # print("seq_dic[user_seq]: ", seq_dic['user_seq'])
    if(len(seq_dic['user_seq'])==0):
        return 0
    args.item_size = int(max_item) + 1
    # print("item_size:", args.item_size)

    def find_new_file(dir):
        '''查找目录下最新的文件'''
        file_lists = os.listdir(dir)
        file_lists.sort(key=lambda fn: os.path.getmtime(dir + "\\" + fn)
        if not os.path.isdir(dir + "\\" + fn) else 0)
        # print('最新的文件为： ' + file_lists[-1])
        file = os.path.join(dir, file_lists[-1])
        # print('完整路径：', file)
        return file_lists[-1]

    dir = r'output\\'
    str_ = find_new_file(dir)
    ind = str_.find(".")
    args.load_model = str_[:ind]
    # print("load_model:", args.load_model)

    # save model args
    cur_time = get_local_time()
    if args.no_filters:
        args.model_name = "SASRec"
    args_str = f'{args.model_name}-{args.data_name}-{cur_time}'
    args.log_file = os.path.join(args.output_dir, args_str + '.txt')
    # print(str(args))
    with open(args.log_file, 'a') as f:
        f.write(str(args) + '\n')

    # save model
    args.checkpoint_path = os.path.join(args.output_dir, args_str + '.pt')
    # print(seq_dic)
    train_dataloader, eval_dataloader, test_dataloader = get_dataloder(args, seq_dic)

    model = FMLPRecModel(args=args)
    trainer = FMLPRecTrainer(model, train_dataloader, eval_dataloader,
                             test_dataloader, args)

    if args.full_sort:
        args.valid_rating_matrix, args.test_rating_matrix = get_rating_matrix(args.data_name, seq_dic, max_item)
        # print(args.test_rating_matrix[0])
    if args.do_eval:
        if args.load_model is None:
            print(f"No model input!")
            exit(0)
        else:
            args.checkpoint_path = os.path.join(args.output_dir, args.load_model + '.pt')
            trainer.load(args.checkpoint_path)
            print(f"Load model from {args.checkpoint_path} for test!")
            scores, result_info = trainer.test(0, full_sort=args.full_sort)

    else:
        early_stopping = EarlyStopping(args.checkpoint_path, patience=args.patience, verbose=True)
        for epoch in range(args.epochs):
            trainer.train(epoch)
            scores, _ = trainer.valid(epoch, full_sort=args.full_sort)
            # evaluate on MRR
            early_stopping(np.array(scores[-1:]), trainer.model)
            if early_stopping.early_stop:
                # print("Early stopping")
                break

        # print("---------------Sample 99 results---------------")
        # load the best model
        trainer.model.load_state_dict(torch.load(args.checkpoint_path))
        scores, result_info = trainer.test(0, full_sort=args.full_sort)

    # print(args_str)
    # print(result_info)
    with open(args.log_file, 'a') as f:
        f.write(args_str + '\n')
        f.write(result_info + '\n')


# main2()
