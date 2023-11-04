# str = "INSERT INTO `video` VALUES "
# print(len(str))
import csv
import pandas as pd

datas = []

with open("video.txt", 'r', encoding='utf-8') as f:
    data = []
    for line in f:
        line = line[28:-2]

        # line = line.strip("'")
        if (line == ''):
            continue

        line = line.split(",")
        a = line[3].strip("'").strip(" ")
        b = (line[0].strip("'")).strip(" ")
        c = (line[2].strip("'")).strip(" ").strip("'")

        data.append(a)
        data.append(b)
        data.append('0')
        data.append('0')
        data.append(c)
        datas.append(data)
        data = []

# title = ['uid', 'vid', 'isLike', 'isCollect', 'timestamp']
# try_ = pd.DataFrame(columns=title, data=datas)
# try_[:, 1:].to_csv("video-log.csv", index_label=False)
with open("video-log.csv", "w",newline='') as f:
    writer = csv.writer(f)
    writer.writerow(datas)

# for i,item in enumerate(line):
#     item = item.strip("'")

# print(line)
