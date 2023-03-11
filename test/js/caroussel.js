var currentSlide = 0;
var slides = document.getElementsByClassName("slide");

function showSlide(n) {
  if (n >= slides.length) {
    n = 0;
  } else if (n < 0) {
    n = slides.length - 1;
  }
  for (var i = 0; i < slides.length; i++) {
    slides[i].classList.remove("active");
  }
  slides[n].classList.add("active");
  currentSlide = n;
}

showSlide(currentSlide);

var prevButton = document.createElement("div");
prevButton.classList.add("btn-prev");
prevButton.innerHTML = "&#10094;";
prevButton.addEventListener("click", function() {
  showSlide(currentSlide - 1);
});

var nextButton = document.createElement("div");
nextButton.classList.add("btn-next");
nextButton.innerHTML = "&#10095;";
nextButton.addEventListener("click", function() {
  showSlide(currentSlide + 1);
});

var carousel = document.querySelector(".carousel");
carousel.appendChild(prevButton);
carousel.appendChild(nextButton);
