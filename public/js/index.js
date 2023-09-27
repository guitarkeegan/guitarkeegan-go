const contactEl = document.querySelector("#contact");
const cardEl = document.querySelector("#cardWrapper");
const contactCardEl = document.querySelector("#contactWrapper");
const flipCardInnerEl = document.querySelector(".flip-card-inner");
const arrowWrapperEl = document.querySelector(".arrow-wrapper");


contactEl.addEventListener("click", cardFlip);
arrowWrapperEl.addEventListener("click", cardUnflip);


function cardFlip(){
    contactEl.disabled = true;
    contactEl.style.color = "gray";
    flipCardInnerEl.classList.add("flipper");
}

function cardUnflip(){
    contactEl.disabled = false;
    contactEl.style.color = "#44e929";
    flipCardInnerEl.classList.remove("flipper");
}