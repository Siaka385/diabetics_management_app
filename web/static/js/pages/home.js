const Home = async () => {
  const app = document.getElementById('app');
  if (!app) return;

  app.innerHTML = `
    <!-- Navbar -->
    <header class="sticky top-0 z-50 bg-white shadow-md">
      <nav class="flex justify-between items-center max-w-7xl mx-auto px-6 py-4">
        <div class="text-primary font-bold text-xl">DiaWise</div>
        <div class="hidden md:flex space-x-6">
          <a href="#features" class="hover:text-primary">Features</a>
          <a href="#community" class="hover:text-primary">Community</a>
          <a href="/blog" class="hover:text-primary">Educational Resources</a>
          <a href="/dashboard" class="hover:text-primary">My Dashboard</a>
        </div>
        <button class="md:hidden text-2xl" id="hamburger">☰</button>
      </nav>
      <div class="hidden flex-col space-y-2 px-6 pb-4 md:hidden bg-white shadow" id="mobileMenu">
        <a href="#features" class="hover:text-primary">Features</a>
        <a href="#community" class="hover:text-primary">Community</a>
        <a href="/blog" class="hover:text-primary">Educational Resources</a>
        <a href="/dashboard" class="hover:text-primary">My Dashboard</a>
      </div>
    </header>

    <!-- Hero -->
    <section class="min-h-[80vh] flex items-center bg-gradient-to-br from-background to-accent">
      <div class="max-w-7xl mx-auto grid grid-cols-1 md:grid-cols-2 gap-12 px-6">
        <div class="flex flex-col justify-center">
          <h1 class="text-4xl md:text-5xl font-bold text-primary mb-6">
            Take Control of Your Diabetes Journey, One Step at a Time
          </h1>
          <p class="mb-8 text-lg">
            Join thousands of people who've found their path to better diabetes management
            through personalized support, expert guidance, and a caring community.
          </p>
          <div class="flex space-x-6">
            <a href="/login" class="px-6 py-3 border-2 border-primary text-primary font-semibold rounded-md hover:bg-primary hover:text-white transition">
              Log In
            </a>
            <a href="/signup" class="px-6 py-3 bg-primary text-white font-semibold rounded-md hover:bg-secondary transition">
              Sign Up
            </a>
          </div>
        </div>
        <div class="relative">
          <img src="/static/images/hero-image.jpg" alt="Diabetes management" class="w-full h-[500px] object-cover rounded-lg shadow-lg"/>
        </div>
      </div>
    </section>

    <!-- Features -->
    <section id="features" class="py-24 bg-white">
      <div class="max-w-7xl mx-auto px-6 text-center">
        <h2 class="text-3xl font-bold mb-12">Your Complete Diabetes Management Companion</h2>
        <div class="grid gap-8 md:grid-cols-3">
          <div class="p-8 bg-background border-2 border-primary rounded-xl shadow hover:-translate-y-2 transition transform">
            <svg class="w-16 h-16 mx-auto mb-4 text-primary" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"/>
            </svg>
            <h3 class="text-xl font-semibold mb-2">Expert-Backed Learning</h3>
            <p>Access clinically-reviewed diabetes education, from blood sugar management to lifestyle strategies.</p>
          </div>
          <div class="p-8 bg-background border-2 border-primary rounded-xl shadow hover:-translate-y-2 transition transform">
            <svg class="w-16 h-16 mx-auto mb-4 text-primary" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
            </svg>
            <h3 class="text-xl font-semibold mb-2">Never Walk Alone</h3>
            <p>Join supportive group discussions led by certified diabetes educators. Share experiences and learn together.</p>
          </div>
          <div class="p-8 bg-background border-2 border-primary rounded-xl shadow hover:-translate-y-2 transition transform">
            <svg class="w-16 h-16 mx-auto mb-4 text-primary" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4"/>
            </svg>
            <h3 class="text-xl font-semibold mb-2">Practical Daily Guidance</h3>
            <p>Get personalized recommendations for meals, exercise, and lifestyle choices that fit your needs.</p>
          </div>
        </div>
      </div>
    </section>

    <!-- CTA -->
    <section id="community" class="py-24 text-center bg-gradient-to-br from-primary to-secondary text-white">
      <div class="max-w-4xl mx-auto px-6">
        <h2 class="text-3xl font-bold mb-6">Your Success Story Starts Here</h2>
        <p class="mb-8 text-lg">Join a community where thousands of people are achieving their diabetes management goals through shared experiences and mutual support.</p>
        <a href="/signup" class="px-6 py-3 border-2 border-white text-white rounded-md hover:bg-white hover:text-primary transition">
          Start Your Journey Today
        </a>
      </div>
    </section>
  `;

  // Hamburger toggle
  const hamburger = document.getElementById('hamburger');
  const menu = document.getElementById('mobileMenu');
  if (hamburger && menu) {
    hamburger.addEventListener('click', () => {
      menu.classList.toggle('hidden');
      menu.classList.toggle('flex');
    });
  }
};

export default Home;
