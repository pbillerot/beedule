/**
 * CHART.JS
 * parcourir les classes pour activer le script de creation des canvas
 */
$(document).ready(function () {
    // Fonctions jouées juste après le chargement de la page html
    $(".bee-chart-quotes").each(function () {
        drawChart($(this));
    });
});
/**
 * PARAMETRES data du canvas :
 * title: 
 * quotes: 
 * quotep
 * min: 
 * max: 
 * minp: 
 * maxp:
 * seuilv: 
 * seuilr: 
 * labels:
 * 
 * @param {*} canvas 
 */
function drawChart(canvas) {
    // Calcul préalable des datasets
    var min = parseFloat(canvas.data('min'));
    var max = parseFloat(canvas.data('max'));
    var fseuilr = parseFloat(canvas.data('seuilr'));
    var fseuilv = parseFloat(canvas.data('seuilv'));
    if (fseuilv > max) {
        max = fseuilv
    }
    if (fseuilr < min) {
        min = fseuilr
    }

    var minp = canvas.data('minp').split(',');
    var maxp = canvas.data('maxp').split(',');
    var barHigh = new Array(minp.length);
    var barLow = new Array(minp.length);
    var zerop = new Array(minp.length);
    var seuilr = new Array(minp.length);
    var seuilv = new Array(minp.length);
    var barMin = 200.0
    var barMax = -200.0
    for (i = 0; i < minp.length; i++) {
        if (parseFloat(minp[i]) < barMin) {
            barMin = parseFloat(minp[i])
        }
        if (parseFloat(maxp[i]) > barMax) {
            barMax = parseFloat(maxp[i])
        }
        var bHigh = new Array(2);
        var bLow = new Array(2);
        if (minp[i] <= 0 && maxp[i] <= 0) {
            // rouge seul
            bLow[0] = minp[i]
            bLow[1] = maxp[i]
            bHigh[0] = null
            bHigh[1] = null
        } else if (minp[i] < 0 && maxp[i] > 0) {
            // rouge et vert
            bLow[0] = minp[i]
            bLow[1] = 0
            bHigh[0] = 0
            bHigh[1] = maxp[i]
        } else {
            // vert seul
            bLow[0] = null
            bLow[1] = null
            bHigh[0] = minp[i]
            bHigh[1] = maxp[i]
        }
        barHigh[i] = bHigh;
        barLow[i] = bLow;
        seuilr[i] = fseuilr;
        seuilv[i] = fseuilv;
        zerop[i] = 0
    }
    // Création du dessin
    var myChart = new Chart(canvas, {
        type: 'bar',
        data: {
            labels: canvas.data('labels').split(','),
            datasets: [
                {
                    type: 'line',
                    label: 'Cotation',
                    borderColor: 'purple',
                    borderWidth: 3,
                    fill: false,
                    data: canvas.data('quotes').split(','),
                    yAxisID: 'cotation'
                },
                {
                    type: 'line',
                    label: 'Seuil Vente',
                    borderColor: 'purple',
                    borderWidth: 2,
                    pointRadius: 0,
                    pointHoverRadius: 0,
                    fill: false,
                    data: seuilv,
                    yAxisID: 'cotation'
                },
                {
                    type: 'line',
                    label: 'Seuil Rentabilité',
                    borderColor: 'purple',
                    borderWidth: 3,
                    borderDash: [5, 5],
                    pointRadius: 0,
                    pointHoverRadius: 0,
                    fill: false,
                    data: seuilr,
                    yAxisID: 'cotation'
                },
                {
                    type: 'line',
                    label: 'Cote %',
                    borderColor: 'grey',
                    borderWidth: 2,
                    borderDash: [5, 5],
                    fill: false,
                    data: canvas.data('quotep').split(','),
                    yAxisID: 'pourcentage'
                },
                {
                    type: 'line',
                    label: 'ligne zéro %',
                    borderColor: 'darkgrey',
                    borderWidth: 2,
                    borderDash: [5, 5],
                    fill: false,
                    data: zerop,
                    pointRadius: 0,
                    pointHoverRadius: 0,
                    yAxisID: 'pourcentage'
                },
                {
                    label: 'Positif %',
                    data: barHigh,
                    backgroundColor: '#c2f0c2',
                    yAxisID: 'pourcentage'
                },
                {
                    label: 'Négatif %',
                    data: barLow,
                    backgroundColor: 'mistyrose',
                    yAxisID: 'pourcentage'
                },
            ]
        },
        options: {
            plugins: {
                title: {
                    display: true,
                    text: canvas.data('title'),
                    font: {
                        style: 'bold',
                        size: 14
                    },
                    padding: {
                        top: 10,
                        bottom: 30
                    }
                }
            },
            scales: {
                cotation: {
                    display: true,
                    min: min - min * 0.01,
                    max: max + max * 0.01,
                    scaleLabel: {
                        display: true,
                        labelString: 'Cotation en €',
                        font: {
                            style: 'bold',
                        }
                    },
                    gridLines: {
                        display: false,
                    },
                    position: 'left'
                },
                pourcentage: {
                    display: true,
                    min: barMin + barMin * 0.1,
                    max: barMax + barMax * 0.1,
                    scaleLabel: {
                        display: true,
                        labelString: 'Évolution en %',
                        font: {
                            style: 'bold'
                        }
                    },
                    position: 'right'
                },
                x: {
                    stacked: true,
                    display: true,
                    scaleLabel: {
                        display: true,
                        labelString: 'Jour de cotation',
                        font: {
                            style: 'bold'
                        }
                    },
                }
            }
        }
    });
}