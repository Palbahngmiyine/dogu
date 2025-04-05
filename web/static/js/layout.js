// 폰트 로딩 완료 시 클래스 추가
document.fonts.ready.then(function () {
    document.documentElement.classList.add('fonts-loaded');
});

// 테마 관련 함수
function setTheme(theme) {
    document.documentElement.setAttribute('data-theme', theme);
    localStorage.setItem('theme', theme);
    updateThemeIcon();
}

function toggleTheme() {
    const currentTheme = document.documentElement.getAttribute('data-theme');
    const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
    setTheme(newTheme);
}

// 테마 아이콘 업데이트
function updateThemeIcon() {
    const themeIcon = document.getElementById('theme-icon');
    if (themeIcon) {
        const currentTheme = document.documentElement.getAttribute('data-theme');
        themeIcon.textContent = currentTheme === 'dark' ? 'dark_mode' : 'light_mode';
    }
}

// 시스템 테마 감지
function setSystemTheme() {
    if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
        setTheme('dark');
    } else {
        setTheme('light');
    }
}

// Navigation Drawer 관련 함수
function toggleDrawer() {
    const drawer = document.querySelector('.navigation-drawer');
    const overlay = document.querySelector('.navigation-drawer-overlay');
    drawer.classList.toggle('open');
    overlay.classList.toggle('open');
}

// DOM이 로드된 후 초기화
document.addEventListener('DOMContentLoaded', function() {
    // 초기 테마 설정
    const savedTheme = localStorage.getItem('theme');
    if (savedTheme) {
        setTheme(savedTheme);
    } else if (window.matchMedia) {
        setSystemTheme();
        window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', setSystemTheme);
    }

    // 초기 테마 아이콘 설정
    updateThemeIcon();
}); 