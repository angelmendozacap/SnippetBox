let navLinks = document.querySelectorAll("nav a")

let currentNavLink = [...navLinks].find(navLink => navLink.getAttribute('href') === window.location.pathname)

if (currentNavLink) {
  currentNavLink.classList.add('live')
}
