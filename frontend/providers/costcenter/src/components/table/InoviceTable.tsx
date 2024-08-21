import { InvoiceTableHeaders, TableHeaderID } from '@/constants/billing';
import useEnvStore from '@/stores/env';
import { ReqGenInvoice, InvoicePayload, BillingType, BillingItem } from '@/types';
import {
  TableContainer,
  Table,
  Thead,
  Tr,
  Th,
  Flex,
  Tbody,
  Td,
  Checkbox,
  Text,
  Button,
  Circle,
  Badge,
  useStatStyles,
  useStyleConfig
} from '@chakra-ui/react';
import { format, parseISO, subHours } from 'date-fns';
import { useTranslation } from 'next-i18next';
import CurrencySymbol from '../CurrencySymbol';
import { BaseTable } from './InoviceBaseTable';
import {
  createColumnHelper,
  HeaderContext,
  CellContext,
  useReactTable,
  getCoreRowModel
} from '@tanstack/react-table';
import { useMemo } from 'react';
import InvoiceDetails from './InvoiceDetails';
import { formatMoney } from '@/utils/format';
import { InvoiceStatus } from '../invoice/Status';
export function InvoiceTable({
  data,
  onSelect,
  toInvoiceDetail
}: {
  toInvoiceDetail: () => void;
  data: InvoicePayload[];
  onSelect?: (type: boolean, item: InvoicePayload) => void;
}) {
  const { t } = useTranslation();
  const needSelect = !!onSelect;
  const currency = useEnvStore((s) => s.currency);
  const columns = useMemo(() => {
    const columnHelper = createColumnHelper<InvoicePayload>();
    const customTh = (needCurrency?: boolean) =>
      function CustomTh({ header }: HeaderContext<InvoicePayload, unknown>) {
        return (
          <Flex display={'flex'} alignItems={'center'}>
            <Text mr="4px">{t(header.id)}</Text>
            {!!needCurrency && (
              <Text>
                <CurrencySymbol type={currency} />
              </Text>
            )}
          </Flex>
        );
      };
    return [
      columnHelper.accessor((row) => row.updatedAt, {
        id: TableHeaderID.InvoiceCreateTime,
        header: customTh(),
        cell(props) {
          const time = props.cell.getValue();
          return format(new Date(time), 'MM-dd HH:mm');
        }
      }),
      columnHelper.accessor((row) => row.status, {
        id: TableHeaderID.Status,
        header: customTh(),
        cell(props) {
          const status = props.cell.getValue();
          return <InvoiceStatus status={status}></InvoiceStatus>;
        }
      }),
      columnHelper.accessor((row) => row.updatedAt, {
        id: TableHeaderID.InvoiceUpdateTime,
        header: customTh(),
        cell(props) {
          const time = props.cell.getValue();
          const isFinish = props.row.original.status === 'COMPLETED';
          if (!isFinish) return '-';
          return format(new Date(time), 'MM-dd HH:mm');
        }
      }),
      columnHelper.accessor((row) => row.totalAmount, {
        id: TableHeaderID.TotalAmount,
        header: customTh(true),
        cell(props) {
          const amount = props.cell.getValue();
          return (
            <Text gap={'6px'} color={'brightBlue.600'}>
              {formatMoney(amount)}
            </Text>
          );
        }
      }),
      columnHelper.accessor((row) => row.detail, {
        id: TableHeaderID.Handle,
        header: customTh(),
        cell(props) {
          return (
            <InvoiceDetails
              toInvoiceDetail={toInvoiceDetail}
              invoice={props.row.original}
            ></InvoiceDetails>
          );
        },
        enablePinning: true
      })
    ];
  }, [t, currency, needSelect, toInvoiceDetail]);
  const table = useReactTable({
    data,
    state: {
      columnPinning: {
        left: [TableHeaderID.APPName],
        right: [TableHeaderID.Handle]
      }
    },
    columns,
    getCoreRowModel: getCoreRowModel()
  });
  return <BaseTable table={table}></BaseTable>;
}
