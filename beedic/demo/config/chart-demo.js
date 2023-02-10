// Graphique de démonstration
// https://www.chartjs.org/docs/latest/

// Boucle de traitement pour dessiner les graphiques de la page
var elements = document.querySelectorAll('.chart-demo');
elements.forEach(function (element) {
  drawChart(element);
});

function drawChart(element) {
  const myChart = new Chart(element.children.item(0).getContext('2d'), {
    type: 'polarArea',
    data: {
      labels: ['Red', 'Green', 'Yellow', 'Grey', 'Blue'],
      datasets: [{
        label: element.dataset.title,
        data: [11, 16, 7, 3, 14],
        backgroundColor: [
          'rgb(255, 99, 132)',
          'rgb(75, 192, 192)',
          'rgb(255, 205, 86)',
          'rgb(201, 203, 207)',
          'rgb(54, 162, 235)'
        ]
      }]
    }
  });
}