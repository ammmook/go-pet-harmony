class SpecialHeader extends HTMLElement {
    connectedCallback() {
        this.innerHTML = `
            <header>
                <link rel="stylesheet" href="static/style_header.css?v=1.0">
                <div class="glass-morphism"></div>
                <div class="header-container">
                    <!-- Logo and Title -->
                    <div class="brand">
                        <a href="/" class="logo-link">
                            <div class="logo-container">
                                <img src="assets/img/logo.png" alt="Pet Harmony Logo">
                                <div class="logo-glow"></div>
                            </div>
                            <h1>Pet Harmony</h1>
                        </a>
                    </div>

                    <!-- Navigation -->
                    <nav class="nav-container">
                        <ul class="main-nav">
                            <li><a href="/Project_PetHotel/" class="nav-link"> <span class="nav-text">Home</span> <span class="nav-indicator"></span></a></li>
                            <li><a href="#service" class="nav-link"> <span class="nav-text">Service</span> <span class="nav-indicator"></span></a></li>
                            <li><a href="#about" class="nav-link"> <span class="nav-text">About</span> <span class="nav-indicator"></span></a></li>
                            <li><a href="contact" class="nav-link"> <span class="nav-text">Contact</span> <span class="nav-indicator"></span></a></li>
                        </ul>

                        <!-- User Navigation -->
                        <div class="user-nav">
                            <div class="user-menu" style="display: none;" id="userMenu">
                                <a class="btn btn-user" onclick="toggleDropdown()">
                                    <img alt="User profile" src="assets/img/user.png"> <span>User</span>
                                </a>

                                <div id="userDropdown" class="dropdown-menu">
                                    <a href="yourprofile" class="menu-item"> <img src="assets/img/user.png" alt="Profile">
                                        <span>Profile</span></a>
                                    <a href="booking" class="menu-item"> <img src="assets/img/appointment.png" alt="Appointment">
                                        <span>Booking</span></a>
                                    <a href="listpets" class="menu-item"> <img src="assets/img/paws.png" alt="Paws">
                                        <span>Pets</span></a>
                                    <a href="logout" class="menu-item" onclick="return confirm('Do you want to logout?')"> <img
                                            src="assets/img/power-on.png" alt="Logout"> <span>Logout</span></a>
                                </div>
                            </div>

                            <div class="auth-buttons" style="display: none;" id="authButtons">
                                <a href="login" class="btn btn-login"> <img src="assets/img/enter.png" alt="Login">
                                    <span>Login</span></a>
                                <a href="register" class="btn btn-signup"> <img src="assets/img/add-user.png" alt="Sign up">
                                    <span>Sign-up</span></a>
                            </div>
                        </div>
                    </nav>
                </div>
            </header>
        `;
    }
}

class SpecialFooter extends HTMLElement {  // Corrected here
    connectedCallback() {
        this.innerHTML = `
            <footer>
            <link rel="stylesheet" href="static/style_footer.css?v=1.0">
                <div class="container">
                    <p>&copy; 2024 Pet Harmony. สงวนลิขสิทธิ์</p>
                </div>
            </footer>
        `;
    }
}

customElements.define('special-header', SpecialHeader);
customElements.define('special-footer', SpecialFooter);

function toggleDropdown() {
    const dropdown = document.getElementById('userDropdown');
    dropdown.style.display = dropdown.style.display === 'block' ? 'none' : 'block';
}
