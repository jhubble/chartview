<template>
    <div class="controls">
        <h1>Status Monitor</h1>
        <v-card flat>
            <v-card-text>
                <v-container fluid>
                    <v-layout row wrap>
                        <v-flex xs12 sm3 md3>
                            <v-checkbox label="cpu"
                                        v-model="cpu"
                                        value="cpu"
                                        hide-details></v-checkbox>
                        </v-flex>
                        <v-flex xs12 sm3 md3>
                            <v-checkbox label="storage"
                                        v-model="storage"
                                        value="storage"
                                        hide-details></v-checkbox>
                        </v-flex>
                        <v-flex xs12 sm3 md3>
                            <v-checkbox label="network"
                                        v-model="network"
                                        value="network"
                                        hide-details></v-checkbox>
                        </v-flex>
                        <v-flex xs12 sm3 md3>
                            <v-checkbox label="memory"
                                        v-model="memory"
                                        value="memory"
                                        hide-details></v-checkbox>
                        </v-flex>
                    </v-layout>
                    <v-radio-group v-model="range" row>
                        <v-radio label="Weekly" value="week"></v-radio>
                        <v-radio label="Hourly" value="hour"></v-radio>
                    </v-radio-group>
                </v-container>
            </v-card-text>
        </v-card>
        <v-layout row wrap>
            <chart v-if="cpu" chartType="cpu" :range="range"></chart>
            <chart v-if="storage" chartType="storage" :range="range"></chart>
            <chart v-if="memory" chartType="memory" :range="range"></chart>
            <chart v-if="network" chartType="network" :range="range"></chart>
        </v-layout>
    </div>
</template>

<script>

    import Vue from 'vue'
    import chart from './Chart.vue';

    Vue.component('chart', chart);
    export default {
        data() {
            return {
                range: "week",
                network: null,
                storage: "storage",
                cpu: null,
                memory: null
            }
        }
    }
</script>
