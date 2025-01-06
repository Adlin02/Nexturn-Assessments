const amountInput = document.getElementById("amount-input");
const descriptionInput = document.getElementById("description-input");
const categoryInput = document.getElementById("category-input");
const expenseTable = document.getElementById("expense-table");

let expenses = JSON.parse(localStorage.getItem("expenses")) || [];
let chart;

function addExpense() {
    const amount = parseFloat(amountInput.value);
    const description = descriptionInput.value.trim();
    const category = categoryInput.value;

    if (!amount || !description) {
        alert("Please fill out all fields!");
        return;
    }

    expenses.push({ amount, description, category });
    renderExpenses();
    resetInputs();
}

function deleteExpense(index) {
    expenses.splice(index, 1);
    renderExpenses();
}

function renderExpenses() {
    expenseTable.innerHTML = "";

    expenses.forEach((expense, index) => {
        const row = document.createElement("tr");
        row.innerHTML = `
            <td>${expense.amount.toFixed(2)}</td>
            <td>${expense.description}</td>
            <td>${expense.category}</td>
            <td><button class="btn btn-danger" onclick="deleteExpense(${index})">Delete</button></td>
        `;
        expenseTable.appendChild(row);
    });

    updateChart();
    saveToLocalStorage();
}


function updateChart() {
    const categoryTotals = expenses.reduce((acc, expense) => {
        acc[expense.category] = (acc[expense.category] || 0) + expense.amount;
        return acc;
    }, {});

    const labels = Object.keys(categoryTotals);
    const data = Object.values(categoryTotals);

    if (chart) chart.destroy();

    const ctx = document.getElementById("expense-chart").getContext("2d");
    chart = new Chart(ctx, {
        type: "pie",
        data: {
            labels: labels,
            datasets: [
                {
                    data: data,
                    backgroundColor: ["#FF6384", "#36A2EB", "#FFCE56", "#4BC0C0"],
                },
            ],
        },
    });

    // Update the summary table
    const summaryTableBody = document.getElementById("summary-table-body");
    summaryTableBody.innerHTML = ""; // Clear existing table rows
    labels.forEach((label, index) => {
        const row = document.createElement("tr");
        row.innerHTML = `
            <td>${label}</td>
            <td>$${data[index].toFixed(2)}</td>
        `;
        summaryTableBody.appendChild(row);
    });
}



function resetInputs() {
    amountInput.value = "";
    descriptionInput.value = "";
    categoryInput.value = "Food";
}

function saveToLocalStorage() {
    localStorage.setItem("expenses", JSON.stringify(expenses));
}

document.getElementById("add-expense-button").addEventListener("click", addExpense);

renderExpenses();
