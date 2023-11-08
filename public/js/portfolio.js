var doEl = document.querySelector(".do");
var reEl = document.querySelector(".re");
var meEl = document.querySelector(".me");
var faEl = document.querySelector(".fa");
var fiEl = document.querySelector(".fi");
var solEl = document.querySelector(".sol");
var playButtonEl = document.querySelector("#playButton");
var done = false;
var beet = new Audio("/public/audio/bagatelle-25_01.mp3");

playButtonEl.addEventListener("click", play);
const unclickedColor = "rgb(200, 231, 195);";
doEl.addEventListener("click", enterPorfolio);

function delay() {
  return new Promise((res, reg) => {
    setTimeout(() => {
      res(1);
    }, 210);
  });
}

async function play() {
  beet.play();
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
  meEl.style.backgroundColor = "";
  reEl.style.backgroundColor = "red";
  console.log(await delay());
  reEl.style.backgroundColor = "";
  done = true;
}

function enterPorfolio() {
  console.log("entering portfolio...");
  console.log(window.location.href);
  if (done) {
    window.location.href = window.location.href + "/pass";
  }
}
