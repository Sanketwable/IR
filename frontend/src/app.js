const express = require('express');
const path = require('path');
const app = express();
const hbs = require('hbs');
const port = process.env.PORT || 5000;

//public static path
const static_path = path.join(__dirname, "../public");
const template_path = path.join(__dirname, "../templates/views");
const partials_path = path.join(__dirname, "../templates/partials");
app.set('view engine', 'hbs');

app.set('views', template_path);
hbs.registerPartials(partials_path);

app.use(express.static(static_path));


//routing
app.get("", (req, res) => {
    res.render('index');
});

app.get("/about", (req, res) => {
    res.render('about');
});

app.get("/weather", (req, res) => {
    res.render('weather');
});

app.get("/byprefix",(req,res)=>{
    res.render('weather1');
});

app.get("/bysuffix",(req,res)=>{
    res.render('weather2');
});

app.get("/bysubstring",(req,res)=>{
    res.render('weather3');
});

app.get("*", (req, res) => {
    res.render('404error', {
        errorMsg: 'Oops! go back'
    });
});

app.listen(port, () => {
    console.log(`listening at port ${port}`);
})