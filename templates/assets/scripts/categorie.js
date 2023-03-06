const progressBar = document.querySelector('.progress-bar');
const progressId = document.getElementById("progress")

document.querySelector(".disk").style.display = "none";
function updateProgress(progress) {
    progressBar.style.width = progress + '%';
    progressBar.setAttribute('aria-valuenow', progress);
}

var progress = 100;
var interval = setInterval(function() {
    progress -= 10.0; // Diminuer de 10%
    updateProgress(progress);
    if (progress <= 0) {
        progressId.style.display = "none";
        document.querySelector(".disk").style.display = "block";
    }
}, 1000);