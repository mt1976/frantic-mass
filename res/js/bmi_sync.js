const bmiUserIdElement = document.getElementById('bmiUserId');
const bmiValueElement = document.getElementById('bmiValue');
const bmiWeightInput = document.getElementById('targetWeight');
const bmiEnrichmentElement = document.getElementById('bmiEnrichment');
const bmiGlyphElement = document.getElementById('bmiGlyph');

function debounce(fn, delay) {
  let timer;
  return function (...args) {
    clearTimeout(timer);
    timer = setTimeout(() => fn.apply(this, args), delay);
  };
}

async function updateBMIFromWeight() {
  const userId = bmiUserIdElement.value.trim();
  const weight = parseFloat(bmiWeightInput.value.trim());
  if (!userId || isNaN(weight) || weight <= 0) return;
  const url = `/bmi/calcBMI/${encodeURIComponent(userId)}/${encodeURIComponent(weight)}`;
  try {
    const res = await fetch(url, { method: 'GET' });
    const data = (await res.json()).data?.attributes || {};
    bmiValueElement.value = data.bmi || '';
    bmiEnrichmentElement.value = data.description || '';
    bmiGlyphElement.textContent = data.glyph || '';
  } catch (err) {
    bmiEnrichmentElement.value = 'Error: ' + err.message;
  }
}

async function updateWeightFromBMI() {
  const userId = bmiUserIdElement.value.trim();
  const bmiValue = parseFloat(bmiValueElement.value.trim());
  if (!userId || isNaN(bmiValue) || bmiValue <= 0) return;
  const url = `/bmi/calcWeight/${encodeURIComponent(userId)}/${encodeURIComponent(bmiValue)}`;
  try {
    const res = await fetch(url, { method: 'GET' });
    const data = (await res.json()).data?.attributes || {};
    bmiWeightInput.value = data.weight || '';
    // Optionally update enrichment/glyph if available
    if (data.description) bmiEnrichmentElement.value = data.description;
    if (data.glyph) bmiGlyphElement.textContent = data.glyph;
  } catch (err) {
    bmiEnrichmentElement.value = 'Error: ' + err.message;
  }
  // Enrich the BMI to make sure the enrichment is correct
  await updateBMIFromWeight();
}

const debouncedUpdateBMIFromWeight = debounce(updateBMIFromWeight, 700);
const debouncedUpdateWeightFromBMI = debounce(updateWeightFromBMI, 700);

bmiWeightInput.addEventListener('input', debouncedUpdateBMIFromWeight);
bmiValueElement.addEventListener('input', debouncedUpdateWeightFromBMI);
bmiUserIdElement.addEventListener('input', () => {
  // When userId changes, update both
  debouncedUpdateBMIFromWeight();
  debouncedUpdateWeightFromBMI();
});
