var circle1El = document.querySelector("#circle1");
var circle2El = document.querySelector("#circle2");
var circle3El = document.querySelector("#circle3");
var circle4El = document.querySelector("#circle4");
var circle5El = document.querySelector("#circle5");

circle1El.addEventListener("click", function () {
    console.log("clicked pocketpr");
    window.location.href = window.location.href.replace("/pass", "/1");
});
