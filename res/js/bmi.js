const userIdInput = document.getElementById('user-id');
const weightInput = document.getElementById('weight');
const responseBox = document.getElementById('response');
const spinner = document.getElementById('spinner');

// Debounce helper
function debounce(fn, delay) {
  let timer;
  return function (...args) {
    clearTimeout(timer);
    timer = setTimeout(() => fn.apply(this, args), delay);
  };
}

// Submit to `/BMI/{userId}/{weight}`
async function submitBMI() {
  const userId = userIdInput.value.trim();
  const weight = parseFloat(weightInput.value.trim());

  if (!userId || isNaN(weight) || weight <= 0) {
    return;
  }

  const url = `/bmi/${encodeURIComponent(userId)}/${encodeURIComponent(weight)}`;

  console.log('Submitting BMI:', url);

  spinner.style.display = 'block';
  responseBox.textContent = '';

  try {
    const res = await fetch(url, { method: 'GET' });

    const contentType = res.headers.get('content-type');
    const data = contentType?.includes('application/json')
      ? await res.json()
      : await res.text();

    responseBox.textContent = typeof data === 'string'
      ? data
      : JSON.stringify(data, null, 2);
  } catch (err) {
    responseBox.textContent = 'Error: ' + err.message;
  } finally {
    spinner.style.display = 'none';
  }
}

// Debounced call
const debouncedSubmitBMI = debounce(submitBMI, 700);

// Watch both fields
userIdInput.addEventListener('input', debouncedSubmitBMI);
weightInput.addEventListener('input', debouncedSubmitBMI);