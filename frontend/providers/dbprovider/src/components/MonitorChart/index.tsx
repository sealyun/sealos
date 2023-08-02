import React, { useEffect, useMemo, useRef } from 'react';
import * as echarts from 'echarts';
import { useGlobalStore } from '@/store/global';
import dayjs from 'dayjs';
import { LineStyleMap } from '@/constants/monitor';
import { Flex } from '@chakra-ui/react';

type MonitorChart = {
  data: {
    xData: string[];
    yData: {
      name: string;
      type: string;
      data: string[];
    }[];
  };
  type?: 'blue' | 'deepBlue' | 'green' | 'purple';
  title: string;
  yAxisLabelFormatter?: (value: number) => string;
};

const MonitorChart = ({ type, data, title, yAxisLabelFormatter }: MonitorChart) => {
  const { screenWidth } = useGlobalStore();
  const chartDom = useRef<HTMLDivElement>(null);
  const myChart = useRef<echarts.ECharts>();

  const option = useMemo(
    () => ({
      tooltip: {
        trigger: 'axis',
        // @ts-ignore
        position: (point, params, dom, rect, size) => {
          let xPos = point[0];
          let yPos = point[1] + 10;
          let chartWidth = size.viewSize[0];
          let chartHeight = size.viewSize[1];
          let tooltipWidth = dom.offsetWidth;
          let tooltipHeight = dom.offsetHeight;

          if (xPos + tooltipWidth > chartWidth) {
            xPos = xPos - tooltipWidth;
          }

          if (xPos < 0) {
            xPos = 0;
          }

          return [xPos, yPos];
        }
      },
      grid: {
        left: '0',
        bottom: '4px',
        top: '10px',
        right: '0',
        containLabel: true
      },
      xAxis: {
        type: 'category',
        offset: 4,
        boundaryGap: false,
        axisLabel: {
          interval: 10
        },
        axisTick: {
          show: false
        },
        data: data?.xData?.map((time) => dayjs(parseFloat(time) * 1000).format('HH:mm'))
      },
      yAxis: {
        type: 'value',
        splitNumber: 2,
        boundaryGap: false,
        axisLabel: {
          formatter: yAxisLabelFormatter
        }
      },
      series: data?.yData?.map((item, index) => {
        return {
          name: item.name,
          data: item.data,
          type: 'line',
          // smooth: true,
          showSymbol: false,
          animationDuration: 300,
          animationEasingUpdate: 'linear',
          areaStyle: {
            color: LineStyleMap[index % LineStyleMap.length].backgroundColor
          },
          lineStyle: {
            width: '1',
            color: LineStyleMap[index % LineStyleMap.length].lineColor
          },
          itemStyle: {
            width: 1.5,
            color: LineStyleMap[index % LineStyleMap.length].lineColor
          },
          emphasis: {
            // highlight
            disabled: true
          }
        };
      })
    }),
    [data?.xData, data?.yData]
  );

  // init chart
  useEffect(() => {
    if (!chartDom.current || myChart?.current?.getOption()) return;
    myChart.current = echarts.init(chartDom.current);
    myChart.current && myChart.current.setOption(option);
  }, [chartDom, option]);

  // data changed, update
  useEffect(() => {
    if (!myChart.current || !myChart?.current?.getOption()) return;
    myChart.current.setOption(option);
  }, [data, option]);

  // resize chart
  useEffect(() => {
    if (!myChart.current || !myChart.current.getOption()) return;
    myChart.current.resize();
  }, [screenWidth]);

  return <Flex ref={chartDom} flex={'1 1 80%'} />;
};

export default MonitorChart;
