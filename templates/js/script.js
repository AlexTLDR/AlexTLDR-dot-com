document.addEventListener('DOMContentLoaded', () => {
  const aboutLink = document.querySelector('a[href="/"]');
  const cvLink = document.querySelector('a[href="/cv"]');
  const portfolioLink = document.querySelector('a[href="/portfolio"]');
  const stuttgartGophersLink = document.querySelector('a[href="/stuttgart-gophers"]');

  if (aboutLink) {
    aboutLink.addEventListener('click', (event) => {
      event.preventDefault(); // Prevent the default link behavior
      console.log('Navigating to /');
      // Navigate to the specified path
      window.location.href = '/';
    });
  }

  if (cvLink) {
    cvLink.addEventListener('click', (event) => {
      event.preventDefault(); // Prevent the default link behavior
      console.log('Navigating to /cv');
      // Navigate to the specified path
      window.location.href = '/cv';
    });
  }

  if (portfolioLink) {
    portfolioLink.addEventListener('click', (event) => {
      event.preventDefault(); // Prevent the default link behavior
      console.log('Navigating to /portfolio');
      // Navigate to the specified path
      window.location.href = '/portfolio';
    });
  }

  if (stuttgartGophersLink) {
    stuttgartGophersLink.addEventListener('click', (event) => {
      event.preventDefault(); // Prevent the default link behavior
      console.log('Navigating to /stuttgart-gophers');
      // Navigate to the specified path
      window.location.href = '/stuttgart-gophers';
    });
  }
});

function openImagePopup(imgId) {
  const img = document.getElementById(imgId);
  const popup = window.open("", "Image Popup", "width=800,height=600");
  popup.document.write(`
      <html>
      <head>
          <style>
              body {
                  margin: 0;
                  padding: 0;
                  display: flex;
                  justify-content: center;
                  align-items: center;
                  height: 100vh;
                  background: linear-gradient(300deg, #007a87, #004d56, #002329);
                  background-size: 180% 180%;
                  animation: gradient-animation 18s ease infinite;
              }
              @keyframes gradient-animation {
                  0% {
                      background-position: 0% 50%;
                  }
                  50% {
                      background-position: 100% 50%;
                  }
                  100% {
                      background-position: 0% 50%;
                  }
              }
              img {
                  max-width: 90%;
                  max-height: 90%;
                  object-fit: contain;
              }
          </style>
      </head>
      <body>
          <img src="${img.src}" alt="Popup Image">
      </body>
      </html>
  `);
}

