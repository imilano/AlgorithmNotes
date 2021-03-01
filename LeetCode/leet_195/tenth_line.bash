#Given a text file file.txt, print just the 10th line of the file.
#
#Example:
#
#Assume that file.txt has the following content:
#
#Line 1
#Line 2
#Line 3
#Line 4
#Line 5
#Line 6
#Line 7
#Line 8
#Line 9
#Line 10
#Your script should output the tenth line, which is:
#
#Line 10
#Note:
#1. If the file contains less than 10 lines, what should you output?
#2. There's at least three different solutions. Try to explore all possibilities.
#

#tail -n 3  file.txt 从文件结尾向上打印3行
#tail -n +3 file.txt 以文件的第三行开始，向下打印所有内容
#head -n 3  file.txt 从头开始打印3行
#head -n -3 file.txt 打印文件除了最后3行的所有内容
#


tail -n +10 file.txt | head -n 1

head -n 10 file.txt | tail -n +10