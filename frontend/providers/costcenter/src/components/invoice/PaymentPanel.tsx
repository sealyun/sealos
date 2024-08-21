import useOverviewStore from '@/stores/overview';
import { useEffect, useMemo, useState } from 'react';
import { ApiResp, RechargeBillingData, RechargeBillingItem, ReqGenInvoice } from '@/types';
import { useQuery, useQueryClient } from '@tanstack/react-query';
import request from '@/service/request';
import { useTranslation } from 'next-i18next';
import { Flex, Heading, HStack, Img, TabPanel, Text, useMediaQuery } from '@chakra-ui/react';
import SwitchPage from '@/components/billing/SwitchPage';
import { InvoicePaymentTable } from '../table/InovicePaymentTable';
import { END_TIME, START_TIME } from '@/constants/payment';

export default function PaymentPanel({
  selectbillings,
  orderID,
  setSelectBillings
}: {
  orderID: string;
  selectbillings: ReqGenInvoice['billings'];
  setSelectBillings: (list: RechargeBillingItem[]) => void;
}) {
  const { t } = useTranslation();
  const [page, setPage] = useState(1);
  const [totalPage, setTotalPage] = useState(1);
  const [totalItem, setTotalItem] = useState(0);
  const [pageSize, setPageSize] = useState(10);

  // const { startTime, endTime } = useOverviewStore()
  const endTime = END_TIME;
  const startTime = START_TIME;
  const body = {
    startTime,
    endTime,
    page,
    pageSize,
    paymentID: orderID
  };
  const { data } = useQuery(['billing', 'invoice', body], () => {
    return request<any, ApiResp<RechargeBillingData>>('/api/billing/rechargeBillingList', {
      data: body,
      method: 'POST'
    });
  });
  useEffect(() => {
    if (!data?.data) {
      return;
    }
    const { total, totalPage } = data.data;
    if (totalPage === 0) {
      // search reset
      setTotalPage(1);
      setTotalItem(1);
    } else {
      setTotalItem(total);
      setTotalPage(totalPage);
    }
    if (totalPage < page) {
      setPage(1);
    }
  }, [data?.data]);
  return (
    <TabPanel p="0" display={'flex'} flexDirection={'column'} flex={'auto'}>
      <InvoicePaymentTable
        selectbillings={selectbillings}
        data={(data?.data?.payments || []).filter((item) => !item.InvoicedAt)}
        setSelectBillings={setSelectBillings}
      ></InvoicePaymentTable>
      <SwitchPage
        marginTop={'16px'}
        mx={'auto'}
        currentPage={page}
        totalPage={totalPage}
        totalItem={totalItem}
        pageSize={pageSize}
        setCurrentPage={function (idx: number): void {
          setPage(idx);
        }}
      ></SwitchPage>
    </TabPanel>
  );
}
