import os
import numpy as np
import pandas as pd
import re

pydir = os.getcwd()
path = pydir + "/txt/"
txtfiles = [i for i in os.listdir(path) if ".txt" in i and "_" in i]


# もしtxtフォルダ内のテキストファイル	
year = ""
month = ""
for oneperson in txtfiles:
	if not year:
		year = oneperson[:2]
	elif year != oneperson[:2]:
		print("txtフォルダ内の年数が統一されていません。同じ年のシフト希望データか確認してください。")
	
	if not month:
		month = oneperson[2:4]
	elif month != oneperson[2:4]:
		print("txtフォルダ内の月が統一されていません。同じ月のシフト希望データか確認してください。")

year = txtfiles[0][:2] # 2020年なら20
month = txtfiles[0][2:4] # 3月なら03
namelist = [i.replace(".txt", "") for i in txtfiles]
namelist = [re.sub('[0-9]*_', '', i) for i in namelist]

dates = {"01":31,
		"02":29,
		"03":31,
		"04":30,
		"05":31,
		"06":30,
		"07":31,
		"08":31,
		"09":30,
		"10":31,
		"11":30,
		"12":31 }

df = pd.DataFrame(columns=namelist,
                  index=list(range(1, dates[month] + 1))+["dates", "hours"])

for onepersonfile, name in zip(txtfiles, namelist):
	f = open(path + onepersonfile)
	shiftlist = f.readlines()
	f.close()

	dates = len(shiftlist)
	df[name]["dates"] = dates
	df[name]["hours"] = 0

	for onedayshift in shiftlist:
		onedayshift = (''.join(onedayshift.splitlines())) # 改行コードを消す
		l = onedayshift.split() # 空白区切りで分けてリストに格納

		# シャープが先頭に付いていたらコメントアウトとみなす
		if l[0].startswith("#"):
			continue
		date = int(l[0])

		time = l[1].split("-")
		starttime = float(time[0])
		finishtime = float(time[1])
		hours = finishtime - starttime
		df[name]["hours"] = df[name]["hours"] + hours

		if starttime%1:
			starttime = str(int(starttime)) + ":" + "30"
		else:
			starttime = str(int(starttime)) + ":" + "00"

		if finishtime%1:
			finishtime = str(int(finishtime)) + ":" + "30"
		else:
			finishtime = str(int(finishtime)) + ":" + "00"

		df[name][date] = starttime + "〜" + finishtime

		# print(date, starttime, "〜", finishtime)

df = df.fillna(" ")
print(df)

df.to_csv(year + "_"+ month + ".csv")

