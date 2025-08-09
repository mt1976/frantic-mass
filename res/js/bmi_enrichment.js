const bmiUserIdElement = document.getElementById('bmiUserId');
const bmiValueElement = document.getElementById('bmiValue');
const bmiEnrichmentElement = document.getElementById('bmiEnrichment');
const bmiGlyphElement = document.getElementById('bmiGlyph');

// Debounce helper
function debounce(fn, delay) {
  let timer;
  return function (...args) {
    clearTimeout(timer);
    timer = setTimeout(() => fn.apply(this, args), delay);
  };
}

// Submit to `/bmi_enrichment/{userId}/{bmiValue}`
async function submitBMIEnrichment() {
  const userId = bmiUserIdElement.value.trim();
  const bmiValue = parseFloat(bmiValueElement.value.trim());
  console.log('Submitting BMI enrichment for user:', userId, 'with BMI value:', bmiValue);
  if (!userId || isNaN(bmiValue) || bmiValue <= 0) {
    return;
  }

  const url = `/bmi/enrichment/${encodeURIComponent(userId)}/${encodeURIComponent(bmiValue)}`;

  console.log('Submitting BMI enrichment:', url);

  bmiEnrichmentElement.textContent = '';

  try {
    const res = await fetch(url, { method: 'GET' });

    const contentType = res.headers.get('content-type');
    const data = contentType?.includes('application/json')
      ? await res.json()
      : await res.text();

    console.log('Response:', data);
    const description = data.data?.attributes?.description;
    console.log('description:', description);

    bmiEnrichmentElement.value = typeof data === 'string'
      ? data
      : description || 'No description available';
    bmiGlyphElement.textContent = data.data?.attributes?.glyph || '🟠';
    console.log('BMI enrichment updated:', bmiEnrichmentElement.value);
  } catch (err) {
    bmiEnrichmentElement.value = 'Error: ' + err.message;
    console.error('Error submitting BMI enrichment:', err);
  } finally {
    console.log('BMI enrichment submission completed');
  }
}

// Debounced call
const debouncedSubmitBMIEnrichment = debounce(submitBMIEnrichment, 700);

// Watch both fields
bmiUserIdElement.addEventListener('input', debouncedSubmitBMIEnrichment);
bmiValueElement.addEventListener('input', debouncedSubmitBMIEnrichment);
