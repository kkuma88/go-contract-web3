async function sendTx() {
  const from = document.getElementById("from").value.trim();
  const to = document.getElementById("to").value.trim();
  const amount = document.getElementById("amount").value.trim();

  const res = await fetch("/api/tx/send", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ from, to, amount })
  });

  const data = await res.json();
  document.getElementById("result").innerText = data.message || data.error;
}


