const test = document.querySelector(".test")
const myForm = document.querySelector("#myForm");

test.addEventListener("click", function() {
    const formData = new FormData(myForm);

    // Envoyez une requête POST
    const data = {
        name: "rock"
    };

    // Envoyez une requête POST
    fetch("/api", {
        method: "POST",
        body: JSON.stringify(data),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
    })
        .then(response => response.json())
        .then(data => {
            // Redirigez l'utilisateur vers une nouvelle page
            window.location.replace("/artistinfo");
        })
        .catch(error => {
            console.error('Error:', error);
        });
});

var slider = new Swiper(".swiper-container", {
    slidesPerView: 'auto',
    spaceBetween: 150,
    centeredSlides: true,
    mousewheel: true
})

slider.on('slideChange', function () {
    TweenMax.to('.slide-text span', .2, {
        y: '-100px',
    })
    TweenMax.to('.slide-number span', .2, {
        x: '-100px',
    })
    TweenMax.to('.swiper-slide-active', .5, {
        scale: .85
    })
})

slider.on('slideChangeTransitionEnd', function () {

    TweenMax.to('.slide-text span', .2, {
        y: 0,
        delay: .5
    })
    TweenMax.to('.slide-text span', 0, {
        y: '100px',
    })

    TweenMax.to('.slide-number span', .2, {
        x: 0,
        delay: .7
    })
    TweenMax.to('.slide-number span', 0, {
        x: '100px',
    })

    TweenMax.to('.swiper-slide-active', .5, {
        scale: 1,
        ease: Power4.easeOut,
    })

    TweenMax.to('.swiper-slide-active .slide-text', 0, {
        autoAlpha: 1
    })
    TweenMax.to('.swiper-slide-active .slide-number', 0, {
        autoAlpha: 1
    })

    TweenMax.to('.swiper-slide-next .slide-text', 0, {
        autoAlpha: 0
    })
    TweenMax.to('.swiper-slide-prev .slide-text', 0, {
        autoAlpha: 0
    })

    TweenMax.to('.swiper-slide-next .slide-number', 0, {
        autoAlpha: 0
    })
    TweenMax.to('.swiper-slide-prev .slide-number', 0, {
        autoAlpha: 0
    })
})

TweenMax.to('.swiper-slide-next .slide-text', 0, {
    autoAlpha: 0
})
TweenMax.to('.swiper-slide-prev .slide-text', 0, {
    autoAlpha: 0
})

TweenMax.to('.swiper-slide-next .slide-number', 0, {
    autoAlpha: 0
})
TweenMax.to('.swiper-slide-prev .slide-number', 0, {
    autoAlpha: 0
})

TweenMax.to('.swiper-slide', 0, {
    scale: .85,
})

TweenMax.to('.swiper-slide-active', 0, {
    scale: 1,
})