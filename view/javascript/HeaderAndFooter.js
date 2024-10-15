class SpecialFooter extends HTMLElement {
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
