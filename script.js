const berryTypes = [
  { name: "Bluecrop", family: "北部ハイブッシュ", rarity: "common", reward: 3, color: "#5d70db", shadow: "#2f3c8c" },
  { name: "Duke", family: "北部ハイブッシュ", rarity: "common", reward: 4, color: "#7488ef", shadow: "#3a4ab0" },
  { name: "Spartan", family: "北部ハイブッシュ", rarity: "common", reward: 4, color: "#6b7ce2", shadow: "#32419b" },
  { name: "Patriot", family: "北部ハイブッシュ", rarity: "common", reward: 4, color: "#6f80d7", shadow: "#37448d" },
  { name: "Elliott", family: "北部ハイブッシュ", rarity: "common", reward: 5, color: "#7985d8", shadow: "#3f4a96" },
  { name: "Legacy", family: "北部ハイブッシュ", rarity: "common", reward: 5, color: "#6370c8", shadow: "#2c387e" },
  { name: "Chandler", family: "北部ハイブッシュ", rarity: "rare", reward: 7, color: "#8a7bf2", shadow: "#5642be" },
  { name: "Liberty", family: "北部ハイブッシュ", rarity: "common", reward: 5, color: "#5f75d2", shadow: "#30458c" },
  { name: "Aurora", family: "北部ハイブッシュ", rarity: "rare", reward: 7, color: "#9a83ff", shadow: "#614bc4" },
  { name: "Jersey", family: "北部ハイブッシュ", rarity: "common", reward: 4, color: "#6071bf", shadow: "#31407b" },
  { name: "Berkeley", family: "北部ハイブッシュ", rarity: "common", reward: 4, color: "#6d82d9", shadow: "#3650a3" },
  { name: "Toro", family: "北部ハイブッシュ", rarity: "common", reward: 5, color: "#7687f0", shadow: "#3f52b0" },
  { name: "Northland", family: "ハーフハイブッシュ", rarity: "common", reward: 4, color: "#5b6bc6", shadow: "#2f397f" },
  { name: "Northblue", family: "ハーフハイブッシュ", rarity: "rare", reward: 6, color: "#816df0", shadow: "#4d38ba" },
  { name: "Northcountry", family: "ハーフハイブッシュ", rarity: "common", reward: 4, color: "#6773c5", shadow: "#374183" },
  { name: "Sunshine Blue", family: "サザンハイブッシュ", rarity: "rare", reward: 7, color: "#8d7be9", shadow: "#5749b5" },
  { name: "Misty", family: "サザンハイブッシュ", rarity: "common", reward: 5, color: "#7384de", shadow: "#3a4ca4" },
  { name: "O'Neal", family: "サザンハイブッシュ", rarity: "rare", reward: 7, color: "#8a77ee", shadow: "#503eb7" },
  { name: "Sharpblue", family: "サザンハイブッシュ", rarity: "common", reward: 5, color: "#6679d0", shadow: "#33428e" },
  { name: "Biloxi", family: "サザンハイブッシュ", rarity: "rare", reward: 7, color: "#917ce3", shadow: "#5743ae" },
  { name: "Powderblue", family: "ラビットアイ", rarity: "common", reward: 5, color: "#7085dc", shadow: "#4153a0" },
  { name: "Tifblue", family: "ラビットアイ", rarity: "common", reward: 5, color: "#6480d3", shadow: "#364d90" },
  { name: "Brightwell", family: "ラビットアイ", rarity: "common", reward: 5, color: "#5f77cb", shadow: "#35458a" },
  { name: "Climax", family: "ラビットアイ", rarity: "common", reward: 5, color: "#7188da", shadow: "#4253a0" },
  { name: "Premier", family: "ラビットアイ", rarity: "common", reward: 5, color: "#7384ca", shadow: "#445090" },
  { name: "Gardenblue", family: "ラビットアイ", rarity: "rare", reward: 7, color: "#8678d7", shadow: "#5146a4" },
  { name: "Columbus", family: "ラビットアイ", rarity: "common", reward: 6, color: "#7180c6", shadow: "#42508a" },
  { name: "Pink Lemonade", family: "ラビットアイ", rarity: "epic", reward: 12, color: "#ff93b9", shadow: "#bb507b" },
  { name: "Ochlockonee", family: "ラビットアイ", rarity: "rare", reward: 8, color: "#9179e1", shadow: "#5b47b0" },
  { name: "Baldwin", family: "ラビットアイ", rarity: "rare", reward: 8, color: "#8a72d7", shadow: "#553f9e" }
];

const rarityWeights = {
  common: 72,
  rare: 24,
  epic: 4
};

const state = {
  money: 0,
  discovered: [],
  currentBerryIndex: null,
  currentReward: 0,
  farmers: 0
};

const plant = document.getElementById("plant");
const moneyValue = document.getElementById("moneyValue");
const discoveredValue = document.getElementById("discoveredValue");
const farmerValue = document.getElementById("farmerValue");
const autoValue = document.getElementById("autoValue");
const currentBerryText = document.getElementById("currentBerryText");
const notice = document.getElementById("notice");
const catalog = document.getElementById("catalog");
const growButton = document.getElementById("growButton");
const harvestButton = document.getElementById("harvestButton");
const farmerButton = document.getElementById("farmerButton");
const saveButton = document.getElementById("saveButton");
const catalogCards = [];

function rarityLabel(rarity) {
  if (rarity === "epic") return "超レア";
  if (rarity === "rare") return "レア";
  return "通常";
}

function persistGame(showMessage = false) {
  localStorage.setItem("blueberryCollectorSave", JSON.stringify(state));

  if (showMessage) {
    showNotice("進行を保存しました。GitHub Pagesでも同じ端末なら続きから遊べます。");
  }
}

function showNotice(text) {
  notice.textContent = text;
}

function renderPlant() {
  plant.querySelectorAll(".berry").forEach((node) => node.remove());

  if (state.currentBerryIndex === null) {
    return;
  }

  const berry = berryTypes[state.currentBerryIndex];
  for (let i = 1; i <= 5; i += 1) {
    const node = document.createElement("div");
    node.className = `berry berry-${i}`;
    node.style.setProperty("--berry-color", berry.color);
    node.style.setProperty("--berry-shadow", berry.shadow);
    plant.appendChild(node);
  }
}

function buildCatalog() {
  berryTypes.forEach((berry, index) => {
    const card = document.createElement("article");
    catalogCards[index] = card;
    catalog.appendChild(card);
    updateCatalogCard(index);
  });
}

function updateCatalogCard(index) {
  const berry = berryTypes[index];
  const discovered = state.discovered.includes(index);
  const card = catalogCards[index];

  if (!card) return;

  card.className = `berry-card${discovered ? "" : " locked"}`;

  if (discovered) {
    card.innerHTML = `
      <div class="berry-top">
        <h3 class="berry-name">${berry.name}</h3>
        <span class="badge ${berry.rarity}">${rarityLabel(berry.rarity)}</span>
      </div>
      <div class="berry-meta">
        系統: ${berry.family}<br>
        収穫額: ${berry.reward}
      </div>
    `;
    return;
  }

  card.innerHTML = `
    <div class="berry-top">
      <h3 class="berry-name">???</h3>
      <span class="locked-label">LOCKED</span>
    </div>
    <div class="berry-meta">新しい品種を見つけると解放されます。</div>
  `;
}

function pickBerryIndex() {
  const weightedPool = berryTypes.map((berry, index) => ({
    index,
    weight: rarityWeights[berry.rarity] || 1
  }));
  const totalWeight = weightedPool.reduce((sum, item) => sum + item.weight, 0);
  let roll = Math.random() * totalWeight;

  for (const item of weightedPool) {
    roll -= item.weight;
    if (roll <= 0) {
      return item.index;
    }
  }

  return weightedPool[weightedPool.length - 1].index;
}

function updateStats() {
  moneyValue.textContent = state.money;
  discoveredValue.textContent = `${state.discovered.length} / ${berryTypes.length}`;
  farmerValue.textContent = state.farmers;
  autoValue.textContent = `${state.farmers} / 秒`;
  farmerButton.disabled = state.money < 20;
}

function updateCurrentBerryText() {
  if (state.currentBerryIndex === null) {
    currentBerryText.textContent = "まだ何も育っていません。種をまいて、新しい品種を見つけましょう。";
    return;
  }

  const berry = berryTypes[state.currentBerryIndex];
  currentBerryText.textContent = `${berry.name} (${berry.family}) が実りました。収穫すると ${state.currentReward} コイン獲得できます。`;
}

function loadGame() {
  const raw = localStorage.getItem("blueberryCollectorSave");
  if (!raw) return;

  try {
    const parsed = JSON.parse(raw);
    state.money = parsed.money || 0;
    state.discovered = Array.isArray(parsed.discovered) ? parsed.discovered : [];
    state.currentBerryIndex = Number.isInteger(parsed.currentBerryIndex) ? parsed.currentBerryIndex : null;
    state.currentReward = parsed.currentReward || 0;
    state.farmers = parsed.farmers || 0;
  } catch (error) {
    console.error("save load failed", error);
  }
}

function refresh() {
  renderPlant();
  updateStats();
  updateCurrentBerryText();
  harvestButton.disabled = state.currentBerryIndex === null;
}

growButton.addEventListener("click", () => {
  const index = pickBerryIndex();
  const berry = berryTypes[index];

  state.currentBerryIndex = index;
  state.currentReward = berry.reward;

  if (!state.discovered.includes(index)) {
    state.discovered.push(index);
    updateCatalogCard(index);
    showNotice(`新品種発見: ${berry.name} を図鑑に登録しました。`);
  } else {
    showNotice(`${berry.name} が育ちました。収穫してコインを増やしましょう。`);
  }

  refresh();
  persistGame();
});

harvestButton.addEventListener("click", () => {
  if (state.currentBerryIndex === null) return;

  const berry = berryTypes[state.currentBerryIndex];
  state.money += state.currentReward;
  state.currentBerryIndex = null;
  state.currentReward = 0;
  showNotice(`${berry.name} を収穫して ${berry.reward} コイン獲得しました。`);
  refresh();
  persistGame();
});

farmerButton.addEventListener("click", () => {
  if (state.money < 20) {
    showNotice("農夫を雇うには 20 コイン必要です。");
    return;
  }

  state.money -= 20;
  state.farmers += 1;
  showNotice(`農夫を1人雇いました。これで毎秒 ${state.farmers} コイン入ります。`);
  refresh();
  persistGame();
});

saveButton.addEventListener("click", () => {
  persistGame(true);
});

setInterval(() => {
  if (state.farmers <= 0) return;
  state.money += state.farmers;
  updateStats();
  persistGame();
}, 1000);

window.addEventListener("beforeunload", () => {
  persistGame();
});

loadGame();
buildCatalog();
refresh();
