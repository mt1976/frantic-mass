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
    console.log('glyph:', data.data.attributes.glyph);
    console.log('bmi:', data.data.attributes.bmi);

    console.log('b4 bmiEnrichment:', bmiEnrichmentElement);
    console.log('b4 bmiGlyph:', bmiGlyphElement);
    console.log('b4 bmiValue:', bmiValueElement);

    bmiEnrichmentElement.value = typeof data === 'string'
      ? data.data.attributes.description
      : description || 'No description available';
    bmiGlyphElement.textContent = data.data.attributes.glyph || '🟠';
    bmiValueElement.value = data.data.attributes.bmi || '000.0000';
    bmiEnrichmentElement.textContent = data.data.attributes.description || 'No enrichment available';
    bmiEnrichmentElement.value = data.data.attributes.description || 'No enrichment available';
      console.log('BMI enrichment updated:', bmiEnrichmentElement.value);
     console.log('BMI textContent updated:', bmiEnrichmentElement.textContent);

    console.log('af bmiEnrichment:', bmiEnrichmentElement);
    console.log('af bmiGlyph:', bmiGlyphElement);
    console.log('af bmiValue:', bmiValueElement);

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