import express from "express";
import http from "http";
import { News } from "./news/index";

const port = 3030;
const news = new News();

const app = express();

const fetchNews = (req: any, res: any) => {
  const random = Math.random();
  // return res.send(news.content(JSON.parse(`{ "a" : "${random}"}`)));
  // return res.send(news.content(JSON.parse(`{}`)));
  const url = "http://cat-fact.herokuapp.com/facts/591d9bb2227c1a0020d26826"
  http.get(url, (resp) => {
    let data = "";
    resp.on('data', (chunk) => {
      data += chunk;
    });
    resp.on('end', () => {
      res.send(news.content(JSON.parse(data)));
    });
  }).on("error", (err) => {
    console.log("Error: " + err.message);
  });
}

app.get('/news', fetchNews);

app.listen(port, () => console.log(`Express ON! listening on port ${port}!`));
