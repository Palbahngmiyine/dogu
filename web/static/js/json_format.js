// JSON 관련 함수들
function formatJSON() {
    const input = document.getElementById('jsonInput');
    const errorMsg = document.getElementById('errorMessage');
    try {
        const json = JSON.parse(input.value);
        input.value = JSON.stringify(json, null, 2);
        errorMsg.classList.add('hidden');
    } catch (e) {
        errorMsg.textContent = '유효하지 않은 JSON 형식입니다.';
        errorMsg.classList.remove('hidden');
    }
}

function minifyJSON() {
    const input = document.getElementById('jsonInput');
    const errorMsg = document.getElementById('errorMessage');
    try {
        const json = JSON.parse(input.value);
        input.value = JSON.stringify(json);
        errorMsg.classList.add('hidden');
    } catch (e) {
        errorMsg.textContent = '유효하지 않은 JSON 형식입니다.';
        errorMsg.classList.remove('hidden');
    }
}

function copyJSON() {
    const input = document.getElementById('jsonInput');
    input.select();
    document.execCommand('copy');
}

function clearJSON() {
    const input = document.getElementById('jsonInput');
    const errorMsg = document.getElementById('errorMessage');
    input.value = '';
    errorMsg.classList.add('hidden');
} 