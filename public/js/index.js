const contactEl = document.querySelector("#contact");
const cardEl = document.querySelector("#cardWrapper");
const contactCardEl = document.querySelector("#contactWrapper");
const flipCardInnerEl = document.querySelector(".flip-card-inner");
const arrowWrapperEl = document.querySelector(".arrow-wrapper");
const contactFormEl = document.querySelector("#contact-form");
const emailEl = document.querySelector("#email");
const messageEl = document.querySelector("#message");


contactEl.addEventListener("click", cardFlip);
arrowWrapperEl.addEventListener("click", cardUnflip);
contactFormEl.addEventListener("submit", sendMessage);


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

function sendMessage(event){
    event.preventDefault();

    fetch("/api/contact", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            email: (emailEl.value).trim(),
            message: (messageEl.value).trim(),
        })
    })
    .then(res => console.log(res))
    .then(data => console.log(data))
    .catch(err => console.error(err));

    emailEl.value = "";
    messageEl.value = "";

    return
}