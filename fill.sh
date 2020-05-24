#!/bin/bash
echo "INSERT INTO packs (pack_id, author,name) VALUES (1,1,'pack1');"
for (( i=1; i<48; i++ )) 
do
echo "INSERT INTO stickers(stick_id, pack_id, sticker_link) VALUES ($i, 1, ' https://tpvk.hb.bizmrg.com/uploads/stickers/$i.png');"
echo "INSERT INTO packstickers(stick_id, pack_id) VALUES ($i, 1);"
done
echo "INSERT INTO packs (pack_id, author,name) VALUES (2,1,'pack2');"
for (( i=48; i<62; i++ ))
do
echo "INSERT INTO stickers(stick_id, pack_id, sticker_link) VALUES ($i, 2, ' https://tpvk.hb.bizmrg.com/uploads/stickers/$i.png');"
echo "INSERT INTO packstickers(stick_id, pack_id) VALUES ($i, 2);"
done