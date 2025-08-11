const bmiUserIdElement = document.getElementById('bmiUserId');
const bmiValueElement = document.getElementById('bmiValue');
const bmiEnrichmentElement = document.getElementById('bmiEnrichment');
const bmiGlyphElement = document.getElementById('bmiGlyph');
const bmiWeightInput = document.getElementById('bmiWeightInput');

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
  const userId = bmiUserIdElement.value.trim();
  const weight = parseFloat(bmiWeightInput.value.trim());
  console.log('Submitting BMI for user:', userId, 'with weight:', weight);
  if (!userId || isNaN(weight) || weight <= 0) {
    return;
  }

  const url = `/bmi/calcBMI/${encodeURIComponent(userId)}/${encodeURIComponent(weight)}`;

  console.log('Submitting BMI:', url);

  bmiEnrichmentElement.textContent = '';

  try {
    const res = await fetch(url, { method: 'GET' });

    const contentType = res.headers.get('content-type');
    const data = contentType?.includes('application/json')
      ? await res.json()
      : await res.text();

    console.log('Response:', data);
    const description = data.data.attributes.description;
    console.log('description:', description);

    bmiEnrichmentElement.value = typeof data === 'string'
      ? data
      : description || 'No description available';
    bmiGlyphElement.textContent = data.data.attributes.glyph || '🟠';
    bmiValueElement.value = data.data.attributes.bmi || '000.0000';
    console.log('BMI enrichment updated:', bmiEnrichmentElement.value);
  } catch (err) {
    bmiEnrichmentElement.value = 'Error: ' + err.message;
    console.error('Error submitting BMI:', err);
  } finally {
    console.log('BMI submission completed');
  }
}

// Debounced call
const debouncedSubmitBMI = debounce(submitBMI, 700);

// Watch both fields
bmiUserIdElement.addEventListener('input', debouncedSubmitBMI);
bmiWeightInput.addEventListener('input', debouncedSubmitBMI);