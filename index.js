const bell = document.querySelector("#bell");

function play() {
  bell.muted = false;
  bell
    .play()
    .then(() => console.log("Chrome successfully started"))
    .catch((e) => console.log("Error"));
}

setInterval(play, 2000);
