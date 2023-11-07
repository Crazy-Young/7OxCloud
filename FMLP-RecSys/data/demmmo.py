import pandas as pd

list=[[1,2,3],[4,5,6],[7,9,9]]

name=["one","two","three"]

name2=["a","b","c"]

test=pd.DataFrame(columns=name,data=list)
test.drop(test.columns[0],axis = 1)
print(test)

# test.to_csv("testcsv.csv",encoding="gbk")
