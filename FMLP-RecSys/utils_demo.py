import os
import csv
from scipy.sparse import csr_matrix
from torch.utils.data import DataLoader, RandomSampler, SequentialSampler
from datasets import FMLPRecDataset
import numpy as np
import json


def save(dict_, fileName):
    if isinstance(dict_, str):
        dict_ = eval(dict_)
    with open(fileName, 'w', encoding='utf-8') as f:
        # f.write(str(dict))  # 直接这样存储的时候，读取时会报错JSONDecodeError，因为json读取需要双引号{"aa":"BB"},python使用的是单引号{'aa':'bb'}
        str_ = json.dumps(dict_, ensure_ascii=False)  # TODO：dumps 使用单引号''的dict ——> 单引号''变双引号"" + dict变str
        # print(type(str_), str_)
        f.write(str_)


def load(fileName):
    if not os.path.exists(fileName):
        # os.makedirs(fileName)
        with open(fileName, 'w', encoding='utf-8') as f:
            dict_ = dict()
            return dict_
    else:
        with open(fileName, 'r', encoding='utf-8') as f:
            data = f.readline().strip()
            # print("data:", data)
            dict_ = dict()
            if (data == ''):
                return dict_
            dict_ = json.loads(data)
            return dict_


def ID_Index():
    id2index = load("id2index.json")
    index2id = load("index2id.json")
    return id2index, index2id


def generate_rating_matrix_valid(user_seq, num_users, num_items):
    # three lists are used to construct sparse matrix
    row = []
    col = []
    data = []
    for user_id, item_list in enumerate(user_seq):
        for item in item_list[:-2]:  #
            row.append(user_id)
            col.append(item)
            data.append(1)

    row = np.array(row)
    col = np.array(col)
    data = np.array(data)
    rating_matrix = csr_matrix((data, (row, col)), shape=(num_users, num_items))

    return rating_matrix


def generate_rating_matrix_test(user_seq, num_users, num_items):
    # three lists are used to construct sparse matrix
    row = []
    col = []
    data = []
    for user_id, item_list in enumerate(user_seq):
        for item in item_list[:-1]:  #
            row.append(user_id)
            col.append(item)
            data.append(1)

    row = np.array(row)
    col = np.array(col)
    data = np.array(data)
    rating_matrix = csr_matrix((data, (row, col)), shape=(num_users, num_items))

    return rating_matrix


def generate_rating_matrix_test2(users, user_seq, num_users, num_items):
    # three lists are used to construct sparse matrix
    row = []
    col = []
    data = []
    for user_id, item_list in enumerate(user_seq):
        for item in item_list[:-1]:  #
            row.append(user_id)
            col.append(item)
            data.append(1)

    row = np.array(row)
    col = np.array(col)
    data = np.array(data)
    rating_matrix = csr_matrix((data, (row, col)), shape=(num_users, num_items))

    return rating_matrix


def get_rating_matrix(data_name, seq_dic, max_item):
    num_items = max_item + 1
    valid_rating_matrix = generate_rating_matrix_valid(seq_dic['user_seq'], seq_dic['num_users'], num_items)
    test_rating_matrix = generate_rating_matrix_test(seq_dic['user_seq'], seq_dic['num_users'], num_items)
    return valid_rating_matrix, test_rating_matrix


def custom_sort(row):
    return (row[0], row[4])


def get_user_seqs_and_sample(data_file, sample_file):
    users = []
    user_seq = []
    item_set = set()
    id2index, index2id = ID_Index()
    with open(data_file, "r", encoding="utf-8") as csvfile:
        csv_reader = csv.reader(csvfile)
        items = []
        # flag = 1
        # 排序
        csv_reader = list(csv_reader)
        csv_reader[1:].sort(key=custom_sort)
        # print(csv_reader)

        for csv_row in csv_reader[1:]:
            # print(csv_row)
            user = int(csv_row[0])
            item = csv_row[1]
            if (id2index.get(item) == None):
                ind = len(id2index)
                id2index[str(item)] = str(ind)
                index2id[str(ind)] = str(item)
                save(id2index, "id2index.json")
                save(index2id, "index2id.json")
            item = int(id2index[item])
            state = str(csv_row[2]) + str(csv_row[3])
            time = csv_row[4]
            if (len(users) == 0):
                users.append(user)
                items.append(item)

            elif (user != users[-1]):
                users.append(user)
                # print("user_seq1:", user_seq)
                # print("items:", items)
                user_seq.append(items)
                item_set = item_set | set(items)
                # print("user_seq2:", user_seq)
                # item_set.clear()
                items=[]
            items.append(item)

        item_set = item_set | set(items)
        # print("items:", items)
        user_seq.append(items)
        # print("user_seq:", user_seq)
        mymap = dict()
        ready2del = []
        length = 0
        for i,user in enumerate(users):
            mymap[user] = user_seq[i]
            if(len(user_seq[i])<4):
                # user_seq.remove(i)
                ready2del.append(i)
                length = length + 1
                # del(user_seq[i])
                # del(users[i])

        for i in range(length):
            ind = ready2del[length-i-1]
            del (user_seq[ind])
            del (users[ind])

        data_file2 = data_file[:-4] +"-sorted-mapped"+".json"
        with open(data_file2, 'w', encoding='utf-8') as f:
            str_ = json.dumps(mymap, ensure_ascii=False)  # TODO：dumps 使用单引号''的dict ——> 单引号''变双引号"" + dict变str
            # print(type(str_), str_)
            f.write(str_)
    # print("one end")
    max_item = int(max(item_set))
    num_users = len(user_seq)

    sample_seq = []

    return user_seq, max_item, num_users, sample_seq


def get_seq_dic(args):
    args.data_file = args.data_dir + args.data_name + '.csv'
    args.sample_file = args.data_dir + args.data_name + '_sample.csv'
    user_seq, max_item, num_users, sample_seq = \
        get_user_seqs_and_sample(args.data_file, args.sample_file)
    # print(user_seq)
    # print(sample_seq)
    seq_dic = {'user_seq': user_seq, 'num_users': num_users, 'sample_seq': sample_seq}

    return seq_dic, max_item


def get_dataloder(args, seq_dic):
    train_dataset = FMLPRecDataset(args, seq_dic['user_seq'], data_type='train')
    train_sampler = RandomSampler(train_dataset)
    train_dataloader = DataLoader(train_dataset, sampler=train_sampler, batch_size=args.batch_size)

    # eval_dataset = FMLPRecDataset(args, seq_dic['user_seq'], test_neg_items=seq_dic['sample_seq'],
    #                               data_type='valid')
    eval_dataset = FMLPRecDataset(args, seq_dic['user_seq'],
                                  data_type='valid')
    eval_sampler = SequentialSampler(eval_dataset)
    eval_dataloader = DataLoader(eval_dataset, sampler=eval_sampler, batch_size=args.batch_size)

    # test_dataset = FMLPRecDataset(args, seq_dic['user_seq'], test_neg_items=seq_dic['sample_seq'], data_type='test')
    test_dataset = FMLPRecDataset(args, seq_dic['user_seq'], data_type='test')
    test_sampler = SequentialSampler(test_dataset)
    test_dataloader = DataLoader(test_dataset, sampler=test_sampler, batch_size=args.batch_size)
    return train_dataloader, eval_dataloader, test_dataloader


# import argparse
#
# parser = argparse.ArgumentParser()
# parser.add_argument("--data_dir", default="./data/", type=str)
# parser.add_argument("--output_dir", default="output/", type=str)
# parser.add_argument("--data_name", default="demo", type=str)
# parser.add_argument("--do_eval", action="store_true")
# parser.add_argument("--load_model", default=None, type=str)
#
# # model args
# parser.add_argument("--model_name", default="FMLPRec", type=str)
# parser.add_argument("--hidden_size", default=64, type=int, help="hidden size of model")
# parser.add_argument("--num_hidden_layers", default=2, type=int, help="number of filter-enhanced blocks")
# parser.add_argument("--num_attention_heads", default=2, type=int)
# parser.add_argument("--hidden_act", default="gelu", type=str)  # gelu relu
# parser.add_argument("--attention_probs_dropout_prob", default=0.5, type=float)
# parser.add_argument("--hidden_dropout_prob", default=0.5, type=float)
# parser.add_argument("--initializer_range", default=0.02, type=float)
# parser.add_argument("--max_seq_length", default=50, type=int)
# parser.add_argument("--no_filters", action="store_true",
#                     help="if no filters, filter layers transform to self-attention")
#
# # train args
# parser.add_argument("--lr", default=0.001, type=float, help="learning rate of adam")
# parser.add_argument("--batch_size", default=256, type=int, help="number of batch_size")
# parser.add_argument("--epochs", default=200, type=int, help="number of epochs")
# parser.add_argument("--no_cuda", action="store_true")
# parser.add_argument("--log_freq", default=1, type=int, help="per epoch print res")
# parser.add_argument("--full_sort", action="store_true")
# parser.add_argument("--patience", default=10, type=int,
#                     help="how long to wait after last time validation loss improved")
#
# parser.add_argument("--seed", default=42, type=int)
# parser.add_argument("--weight_decay", default=0.0, type=float, help="weight_decay of adam")
# parser.add_argument("--adam_beta1", default=0.9, type=float, help="adam first beta value")
# parser.add_argument("--adam_beta2", default=0.999, type=float, help="adam second beta value")
# parser.add_argument("--gpu_id", default="0", type=str, help="gpu_id")
# parser.add_argument("--variance", default=5, type=float)
#
# args = parser.parse_args()
# seq_dic, max_item = get_seq_dic(args)
# # set_seed(args.seed)
# # check_path(args.output_dir)
