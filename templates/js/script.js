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