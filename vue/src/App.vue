<template>
  <div class="table-wrapper">
    <div class="form-field">
      <label for="base-date">Base Date:</label>
      <input type="date" id="base-date" v-model="baseDate" @change="fetchData" />
    </div>
    <div class="form-field">
      <label for="at-date">At Date:</label>
      <input type="date" id="at-date" v-model="atDate" @change="fetchData" />
    </div>
    <table class="table">
  <thead>
    <tr>
      <th v-for="column in columns" :key="column" @click="sort(column)">
        {{ column }}
      </th>
    </tr>
  </thead>
  <tbody>
    <tr v-for="row in sortedCards" :key="row.name+row.version+row.rarity+row.type+row.base_price+row.status">
      <td>{{ row.name }}</td>
      <td>{{ row.sub }}</td>
      <td>{{ row.version }}</td>
      <td>{{ row.rarity }}</td>
      <td>{{ row.type }}</td>
      <td>{{ row.base_price }}</td>
      <td>{{ row.at_price }}</td>
      <td>{{ row.status }}</td>
    </tr>
  </tbody>
</table>
  </div>
</template>

<style>
  .form-field {
    display: flex;
    flex-direction: column;
    margin-bottom: 16px;
  }

  .form-field label {
    color: #333;
    font-size: 16px;
    margin-bottom: 8px;
  }
  /* .form-field input[type="date"] {
    border: none;
    border-bottom: 2px solid #ff8a80;
    border-radius: 0;
    padding: 8px;
    font-size: 16px;
  } */

  label {
  display: block;
  font-size: 0.8rem;
  margin-bottom: 0.5rem;
  color: #555;
  }

  input[type="date"] {
    width: 100%;
    padding: 0.5rem;
    border: none;
    border-radius: 4px;
    box-shadow: 0 0 0 1px #ddd;
    position: relative;
    font-size: 1rem;
  }

  input[type="date"]:focus {
    outline: none;
    box-shadow: 0 0 0 2px #2196f3;
  }
  input[type=date]::-webkit-calendar-picker-indicator {
      position: absolute;
      width: 100%;
      height: 100%;
      opacity: 0;
  }
  .table {
    width: 100%;
    border-collapse: collapse;
    overflow: hidden;
    box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
    border-radius: 10px;
    margin-bottom: 1rem;
  }

  .table thead tr {
    background-color: #2196f3;
    color: #fff;
    text-align: left;
  }

  .table th,
  .table td {
    padding: 0.75rem;
    vertical-align: top;
    border-top: 1px solid #dee2e6;
  }

  .table th {
    font-weight: 700;
  }

  .table tbody tr:nth-child(even) {
    background-color: #f8f9fa;
  }

  .table tbody tr:hover {
    background-color: #e9ecef;
  }

  .table tbody tr.active-row {
    font-weight: 700;
    background-color: #f5f5f5;
  }

  .table-wrapper {
    max-width: 800px;
    margin: 0 auto;
    padding: 16px;
    background-color: #f5f5f5;
    border-radius: 4px;
  }
</style>

<script setup>
import { ref, onMounted, computed } from 'vue';

const sortDirection = ref("asc");
const sortingColumn = ref("name");
const columns = [
  "name",
  "sub",
  "version",
  "rarity",
  "type",
  "base_price",
  "at_price",
  "status"
];

const cards = ref([]);
const baseDate = ref('');
const atDate = ref('');

onMounted(async () => {
  setDefaultDate()
  await fetchData()
})

const fetchData = async () => {
  try {
    const response = await fetch(`https://ptcwave.nokotools.tokyo/subprice?base=${baseDate.value}&at=${atDate.value}`);
    const data = await response.json();
    cards.value = data.cards;
  } catch (error) {
    console.error(error);
  }
};

const formatDate = (date) => {
  const padMonth = `0${date.getMonth()+1}`.slice(-2);
  const padDate = `0${date.getDate()}`.slice(-2);
  return `${date.getFullYear()}-${padMonth}-${padDate}`;
}

const setDefaultDate = () => {
  const today = new Date();
  const yesterday = new Date(today);
  yesterday.setDate(today.getDate() - 1);

  baseDate.value = formatDate(yesterday);
  atDate.value = formatDate(today);
};

const sort = (column) => {
  if (column === sortingColumn.value) {
    sortDirection.value = sortDirection.value === 'asc' ? 'desc' : 'asc';
  } else {
    sortingColumn.value = column;
    sortDirection.value = 'asc';
  }
}

const sortedCards = computed(() => {
  if (sortingColumn.value) {
    return cards.value.sort((a, b) => {
      const valueA = a[sortingColumn.value];
      const valueB = b[sortingColumn.value];
      if (typeof valueA === 'string') {
        return valueA.localeCompare(valueB) * (sortDirection.value === 'asc' ? -1 : 1);
      } else {
        return (valueA - valueB) * (sortDirection.value === 'asc' ? -1 : 1);
      }
    });
  } 
  
  return cards.value;
})
</script>