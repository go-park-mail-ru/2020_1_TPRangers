#!/bin/bash
echo "INSERT INTO packs (pack_id, author) VALUES (1,1);"
for (( i=1; i<48; i++ )) 
do
echo "INSERT INTO stickers(stick_id, pack_id, sticker_link) VALUES ($i, 1, 'https://social-hub.ru/uploads/stickers/$i.png');"
echo "INSERT INTO packstickers(stick_id, pack_id) VALUES ($i, 1);" 
done
