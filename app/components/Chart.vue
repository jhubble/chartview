<template>
        <v-flex d-flex xs12 sm12 md12 lg6 xl6 child-flex>
            <v-card flat>
                <v-card-text>
                    <h2>{{chartType}}</h2>
                    <div v-if="chartData" class="chart">
                        <h3>{{descriptiveText[chartType]}}</h3>
                        <line-chart :download="true" :min="0" :max="100" :data="chartData"></line-chart>
                    </div>
                    <div v-else class="error">
                        Unable to load chart:
                        <pre>{{chartData}}</pre>
                    </div>
                </v-card-text>
            </v-card>
        </v-flex>
</template>

<script>
    import axios from 'axios';

    export default {
        data() {
            return {
                chartData: null,
                name: "",
                descriptiveText: {
                    "cpu": "CPU Utilization Percentage",
                    "memory": "Memory Usage Percentage. Available Memory: 32 GB",
                    "storage": "Storage capacity percent used. Total Capacity: 64 TB",
                    "network": "Network utilization percent. Total bandwidth: 1 GB/s"
                },
                errors: []
            }
        },
        props: ['chartType', 'range'],
        // Refresh chart data if time range changes
        watch: {
            'range': {
                handler: function (val, oldVal) {
                    if (val != oldVal) {
                        this.getData();
                    }
                },
                deep: true
            }
        },
        methods: {
            getData: function () {
                axios.get("/data?type=" + this.chartType + '&range=' + this.range)
                    .then(response => {
                        if (response.data && response.data.Data) {
                            console.log("RESPONSE:", response);
                            this.chartData = response.data.Data.map((k) => {
                                // convert second timestamp(unix) to milliseconds (JS)
                                return [new Date(k.timestamp*1000), k.Value];
                            });
                            this.name = response.data.type;
                            console.log("CHART DATA:", this.chartData);
                        }
                        else {
                            //this.errors.push("Data not found");
                        }
                    })
                    .catch(e => {
                        // TODO implement better error catching
                        console.log("ERROR:", e);
                        //this.errors.push(e)
                    })
            }

        },
        mounted: function () {
            this.getData();

        }
    }
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    h1, h2 {
        font-weight: normal;
    }

    h2 {
        text-transform: uppercase;
    }

    ul {
        list-style-type: none;
        padding: 0;
    }

    li {
        display: inline-block;
        margin: 0 10px;
    }

    a {
        color: #42b983;
    }
</style>
