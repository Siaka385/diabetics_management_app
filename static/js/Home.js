  // Get all navigation links
  const navLinks = document.querySelectorAll('.sidebar-nav a');
        
  // Get all main content pages
  const pages = {
      'Dashboard': document.querySelector('.welcomepage'),
      'Blood Sugar': document.querySelector('.bloodsugarMonitorpage'),
      'Diet & Nutrition': document.querySelector('.dietAndNutrion'),
      'Medications': document.querySelector('.med'),
      'Education': document.querySelector('.edu'),
      'Support Group': document.querySelector('.commAndSupport')
  };

  // Function to hide all pages
  function hideAllPages() {
      Object.values(pages).forEach(page => {
          if (page) page.classList.add('hide');
      });
  }

  // Function to remove active class from all nav links
  function removeActiveLinks() {
      navLinks.forEach(link => link.classList.remove('active'));
  }

  // Add click event listener to each navigation link
  navLinks.forEach(link => {
      link.addEventListener('click', (e) => {
       //   e.preventDefault();
          
          // Remove active state from all links
          removeActiveLinks();
          
          // Add active state to clicked link
          link.classList.add('active');
          
          // Hide all pages
          hideAllPages();
          
          // Show the corresponding page
          const pageName = link.textContent.trim();
          const currentPage = pages[pageName];
          
          if (currentPage) {
              currentPage.classList.remove('hide');
          } else {
              console.warn(`No page found for: ${pageName}`);
          }
      });
  });

  // Optionally, set the default active page (Dashboard)
  const dashboardLink = document.querySelector('.sidebar-nav a.active');
  if (dashboardLink) {
      const dashboardPage = pages['Dashboard'];
      if (dashboardPage) {
          dashboardPage.classList.remove('hide');
      }
  }