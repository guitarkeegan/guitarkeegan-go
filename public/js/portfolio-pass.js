var circle1El = document.querySelector("#circle1");
var circle2El = document.querySelector("#circle2");
var circle3El = document.querySelector("#circle3");
var circle4El = document.querySelector("#circle4");
var circle5El = document.querySelector("#circle5");
var mainEl = document.querySelector("main");

circle1El.addEventListener("click", function () {
    mainEl.classList.add("circ-trans-1");
    window.location.href = window.location.href.replace("/pass", "/1");
});

circle2El.addEventListener("click", function () {
    console.log("clicked circle 2")
    mainEl.classList.add("circ-trans-1");
    window.location.href = window.location.href.replace("/pass", "/2");
});
