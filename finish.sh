#!/bin/zsh
echo "LOOKING FOR SERVICES ... \n\n"
echo "main\n $(lsof -i:3001 )\n\n"
echo "auth\n $(lsof -i:3080  )\n\n"
echo "chat\n $(lsof -i:3081 )\n\n"
echo "likes\n $(lsof -i:3082)\n\n"
echo "photos\n $(lsof -i:3083 )\n\n"

echo "photos_save\n $(lsof -i:5000 )\n\n"


kill -9 $(lsof -t -i:3001)
kill -9 $(lsof -t -i:3080)
kill -9 $(lsof -t -i:3081)
kill -9 $(lsof -t -i:3082)
kill -9 $(lsof -t -i:3083)

kill -9 $(lsof -t -i:5000)



echo "CHECKING SERVICES ...\n"
echo "main\n $(lsof -i:3001)\n"
echo "auth\n $(lsof -i:3080)\n"
echo "chat\n $(lsof -i:3081)\n"
echo "likes\n $(lsof -i:3082)\n"
echo "photos\n $(lsof -i:3083)\n"
echo "photos_save\n $(lsof -i:5000)\n"




echo "CHECKING SERVICES ...\n\n"
echo "main\n $(lsof -i:3001)\n\n"
echo "auth\n $(lsof -i:3080)\n\n"
echo "chat\n $(lsof -i:3081)\n\n"
echo "likes\n $(lsof -i:3082)\n\n"
echo "photos\n $(lsof -i:3083)\n\n"

