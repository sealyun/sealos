import React, { useEffect, useMemo, useRef } from 'react';
import * as echarts from 'echarts';
import { useGlobalStore } from '@/store/global';
import dayjs from 'dayjs';
import { LineStyleMap } from '@/constants/monitor';
import { Flex, FlexProps, Text } from '@chakra-ui/react';
import MyIcon from '../Icon';

type MonitorChart = FlexProps & {
  data: {
    xData: string[];
    yData: {
      name: string;
      type: string;
      data: number[];
      lineStyleType?: string;
    }[];
  };
  type?: 'blue' | 'deepBlue' | 'green' | 'purple';
  title: string;
  yAxisLabelFormatter?: (value: number) => string;
  yDataFormatter?: (values: number[]) => number[];
  unit?: string;
  isShowLegend?: boolean;
};

const MonitorChart = ({
  type,
  data,
  title,
  yAxisLabelFormatter,
  yDataFormatter,
  unit,
  isShowLegend = true,
  ...props
}: MonitorChart) => {
  const { screenWidth } = useGlobalStore();
  const chartDom = useRef<HTMLDivElement>(null);
  const myChart = useRef<echarts.ECharts>();

  const option = useMemo(
    () => ({
      tooltip: {
        trigger: 'axis',
        formatter: (params: any) => {
          let axisValue = params[0]?.axisValue;
          const content = params
            .map(
              (item: any) =>
                `${item?.marker} ${item?.seriesName}&nbsp; &nbsp;<span style="font-weight: 500">${
                  item?.value
                }${unit ? unit : ''}</span>  <br/>`
            )
            .join('');
          const str = axisValue + '<br/>' + content;
          return str;
        },
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
        left: '4px',
        bottom: '4px',
        top: '10px',
        right: '20px',
        containLabel: true
      },
      xAxis: {
        show: true,
        type: 'category',
        offset: 4,
        boundaryGap: false,
        axisLabel: {
          interval: (index: number, value: string) => {
            const total = data?.xData?.length || 0;
            if (index === 0 || index === total - 1) return false;
            return index % Math.floor(total / 6) === 0;
          },
          textStyle: {
            color: '#667085'
          },
          hideOverlap: true
        },
        axisTick: {
          show: false
        },
        axisLine: {
          show: true,
          lineStyle: {
            color: '#E4E7EC',
            type: 'solid'
          }
        },
        data: data?.xData?.map((time) => dayjs(parseFloat(time) * 1000).format('MM-DD HH:mm'))
      },
      yAxis: {
        type: 'value',
        splitNumber: 2,
        max: 100,
        min: 0,
        boundaryGap: false,
        axisLabel: {
          formatter: yAxisLabelFormatter
        },
        axisLine: {
          show: false
        },
        splitLine: {
          lineStyle: {
            type: 'dashed',
            color: '#E4E7EC'
          }
        }
      },
      series: data?.yData?.map((item, index) => {
        return {
          name: item.name,
          data: item.data,
          type: 'line',
          smooth: true,
          showSymbol: false,
          animationDuration: 300,
          animationEasingUpdate: 'linear',
          areaStyle: {
            color: LineStyleMap[index % LineStyleMap.length].backgroundColor
          },
          lineStyle: {
            width: '1',
            color: LineStyleMap[index % LineStyleMap.length].lineColor,
            type: item?.lineStyleType || 'solid'
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

  useEffect(() => {
    if (!chartDom.current) return;

    if (!myChart.current) {
      myChart.current = echarts.init(chartDom.current);
    } else {
      myChart.current.dispose();
      myChart.current = echarts.init(chartDom.current);
    }

    myChart.current.setOption(option);
  }, [data, option]);

  useEffect(() => {
    return () => {
      if (myChart.current) {
        myChart.current.dispose();
      }
    };
  }, []);

  // resize chart
  useEffect(() => {
    if (!myChart.current || !myChart.current.getOption()) return;
    myChart.current.resize();
  }, [screenWidth]);

  return (
    <Flex position={'relative'} height={'100%'} gap={'25px'}>
      <Flex ref={chartDom} flex={'1 1 80%'} />
      {isShowLegend && (
        <Flex
          justifyContent={'center'}
          alignContent={'center'}
          flexDirection={'column'}
          flex={'1 0 20%'}
          gap={'12px'}
        >
          {data?.yData?.map((item, index) => (
            <Flex key={item?.name + index} alignItems={'center'} w={'fit-content'}>
              <MyIcon
                width={'16px'}
                name="chart"
                color={LineStyleMap[index % LineStyleMap.length].lineColor}
                mr="6px"
              />
              <Text fontSize={'11px'} color={'grayModern.900'} fontWeight={500}>
                {item?.name}
              </Text>
            </Flex>
          ))}
        </Flex>
      )}
    </Flex>
  );
};

export default MonitorChart;
