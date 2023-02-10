/**
 * Graphique du Pluviomètre
 * https://www.chartjs.org/docs/latest/
 */

// Boucle de traitement pour dessiner les graphiques de la page
var elements = document.querySelectorAll('.chart-pluvio');
elements.forEach(function (element) {
  drawChart(element);
});

function drawChart(element) {
  // Données collectées
  const mm2021 = element.dataset.mm2021.split(',');
  const mm2022 = element.dataset.mm2022.split(',');

  // Données calculées
  var vv2021 = new Array(12).fill(0);
  var vv2022 = new Array(12).fill(0);
  var cc2021 = new Array(12).fill(0);
  var cc2022 = new Array(12).fill(0);
  var c2021 = 0, c2022 = 0;
  for (i = 0; i < 12; i++) {
    // Conversion des données collectées
    if (mm2021[i] !== "undefined") {
      vv2021[i] = parseInt(mm2021[i]);
    }
    if (mm2022[i] !== "undefined") {
      vv2022[i] = parseInt(mm2022[i]);
    }
    // Cumul par mois
    c2021 += vv2021[i];
    cc2021[i] = c2021;
    c2022 += vv2022[i];
    cc2022[i] = c2022;
  }

  // Construction du graphique
  const myChart = new Chart(element.children.item(0).getContext('2d'), {
    type: 'bar',
    data: {
      labels: ['Janv.', 'Fév.', 'Mars', 'Avril', 'Mai', 'Juin', 'Juil.', 'Août', 'Sept.', 'Oct.', 'Nove.', 'Déc.'],
      datasets: [
        {
          type: 'bar',
          backgroundColor: 'rgba(54, 162, 235, 0.2)',
          borderColor: 'rgba(54, 162, 235)',
          borderWidth: 1,
          label: '2021',
          data: vv2021,
          yAxisID: 'vvscale'
        },
        {
          type: 'line',
          borderColor: 'rgba(54, 162, 235)',
          borderWidth: 1,
          label: 'Cumul 2021',
          data: cc2021,
          yAxisID: 'ccscale'
        },
        {
          type: 'bar',
          backgroundColor: 'rgba(153, 102, 255, 0.2)',
          borderColor: 'rgba(153, 102, 255)',
          borderWidth: 1,
          label: '2022',
          data: vv2022,
          yAxisID: 'vvscale'
        },
        {
          type: 'line',
          borderColor: 'rgba(153, 102, 255)',
          borderWidth: 1,
          label: 'Cumul 2022',
          data: cc2022,
          yAxisID: 'ccscale'
        }
      ]
    },
    options: {
      scales: {
        vvscale: {
          display: true,
          title: {
            display: true,
            text: 'Mois en mm',
          },
          position: 'left'
        },
        ccscale: {
          display: true,
          title: {
            display: true,
            text: 'Cumul annuel en mm',
          },
          grid: {
            display: false
          },
          position: 'right'
        }
      }
    }
  });
}