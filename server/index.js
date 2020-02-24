const express = require('express');
const path = require('path');
const app = express();
app.set("view engine", "pug");
app.set("views", path.join(__dirname,"..", "public/templates"));
app.use(express.static(path.resolve(__dirname, '..', 'public')));

app.listen(3000, function() {
  console.log('Example app listening on port 3000!');
});
