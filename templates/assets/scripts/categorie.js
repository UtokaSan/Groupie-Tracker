let progress = document.querySelector('.progress');

let intervalId = setInterval(function() {
    let width = parseInt(progress.style.width) || 0;
    width++;
    progress.style.width = width + '%';
    if (width >= 100) {
        clearInterval(intervalId);
    }
}, 100);

setTimeout(function() {
    document.querySelector('.progress').style.width = '0%';
}, 10000);