const progressBar = document.querySelector('.progress-bar');
const progressId = document.getElementById("progress")

function updateProgress(progress) {
    progressBar.style.width = progress + '%';
    progressBar.setAttribute('aria-valuenow', progress);
}

var progress = 100;
var interval = setInterval(function() {
    progress -= 12.8; // Diminuer de 10%
    updateProgress(progress);
    if (progress <= 0) {
        progressId.style.display = "none";
    }
}, 1000);