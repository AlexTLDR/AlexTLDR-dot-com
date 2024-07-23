document.addEventListener('DOMContentLoaded', () => {
    const aboutLink = document.querySelector('a[href="/"]');
    const cvLink = document.querySelector('a[href="/cv"]');
  
    aboutLink.addEventListener('click', (event) => {
      event.preventDefault(); // Prevent the default link behavior
      console.log('Navigating to /');
      // Navigate to the specified path
      window.location.href = '/';
    });
  
    cvLink.addEventListener('click', (event) => {
      event.preventDefault(); // Prevent the default link behavior
      console.log('Navigating to /cv');
      // Navigate to the specified path
      window.location.href = '/cv';
    });
  });
