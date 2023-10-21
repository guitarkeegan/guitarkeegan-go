var mainEl = document.querySelector("main");
mainEl.style.filter = "brightness(.1)";

setTimeout(() => {
  mainEl.style.filter = "brightness(1)";
}, 300);
