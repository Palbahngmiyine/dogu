// 폰트 로딩 완료 시 클래스 추가
document.fonts.ready.then(function () {
    document.documentElement.classList.add('fonts-loaded');
});

// 테마 관련 함수
function setTheme(theme) {
    document.documentElement.setAttribute('data-theme', theme);
    localStorage.setItem('theme', theme);
}

function toggleTheme() {
    const currentTheme = document.documentElement.getAttribute('data-theme');
    const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
    setTheme(newTheme);
}

// 시스템 테마 감지
function setSystemTheme() {
    if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
        setTheme('dark');
    } else {
        setTheme('light');
    }
}

// 초기 테마 설정
const savedTheme = localStorage.getItem('theme');
if (savedTheme) {
    setTheme(savedTheme);
} else if (window.matchMedia) {
    setSystemTheme();
    window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', setSystemTheme);
}

// Navigation Drawer 관련 함수
function toggleDrawer() {
    const drawer = document.querySelector('.navigation-drawer');
    const overlay = document.querySelector('.navigation-drawer-overlay');
    drawer.classList.toggle('open');
    overlay.classList.toggle('open');
} 