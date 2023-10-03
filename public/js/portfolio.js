var doEl = document.querySelector(".do");
var reEl = document.querySelector(".re");
var meEl = document.querySelector(".me");
var faEl = document.querySelector(".fa");
var fiEl = document.querySelector(".fi");
var solEl = document.querySelector(".sol");
var playButtonEl = document.querySelector("#playButton");

var done = false;

playButtonEl.addEventListener("click", play);
const unclickedColor = "rgb(200, 231, 195);"

function delay() {
    return new Promise((res, reg)=>{
        setTimeout(()=>{
            res(1)
        }, 400);
    });
}

async function play() {
    solEl.style.backgroundColor = "red";
    console.log(await delay());
    solEl.style.backgroundColor = "";
    fiEl.style.backgroundColor = "red";
    console.log(await delay());
    fiEl.style.backgroundColor = "";
    solEl.style.backgroundColor = "red";
    console.log(await delay());
    solEl.style.backgroundColor = "";
    fiEl.style.backgroundColor = "red";
    console.log(await delay());
    fiEl.style.backgroundColor = "";
    solEl.style.backgroundColor = "red";
    console.log(await delay());
    solEl.style.backgroundColor = "";
    reEl.style.backgroundColor = "red";
    console.log(await delay());
    reEl.style.backgroundColor = "";
    faEl.style.backgroundColor = "red";
    console.log(await delay());
    faEl.style.backgroundColor = "";
    meEl.style.backgroundColor = "red";
    console.log(await delay());
    done = true;
}

function enterPorfolio(){

}