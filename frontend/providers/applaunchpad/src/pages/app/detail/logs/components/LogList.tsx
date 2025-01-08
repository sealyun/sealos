import {
  createColumnHelper,
  flexRender,
  getCoreRowModel,
  useReactTable
} from '@tanstack/react-table';
import { useState } from 'react';
import { useTranslation } from 'next-i18next';
import {
  Box,
  Button,
  Collapse,
  Divider,
  Flex,
  Text,
  Checkbox,
  CheckboxGroup
} from '@chakra-ui/react';

import MyIcon from '@/components/Icon';

interface FieldItem {
  value: string;
  label: string;
  checked: boolean;
}

export const LogList = () => {
  const { t, i18n } = useTranslation();
  const lang = i18n.language;

  const [onOpenField, setOnOpenField] = useState(false);
  const [hiddenFieldCount, setHiddenFieldCount] = useState(0);
  const [visibleFieldCount, setVisibleFieldCount] = useState(0);

  const [fieldList, setFieldList] = useState<FieldItem[]>([
    {
      value: 'test',
      label: 'test',
      checked: true
    }
  ]);

  return (
    <Flex flexDir={'column'}>
      <Flex pl={4} alignItems={'center'} gap={4} justifyContent={'space-between'}>
        <Flex alignItems={'center'} gap={4}>
          <Text
            bg={'transparent'}
            border={'none'}
            boxShadow={'none'}
            color={'grayModern.900'}
            fontWeight={500}
            fontSize={'14px'}
            lineHeight={'20px'}
          >
            {t('Log')}
          </Text>
          <Flex
            alignItems={'center'}
            bg={'grayModern.50'}
            borderRadius={'8px'}
            {...(onOpenField && {
              borderRadius: '8px 8px 0 0'
            })}
            px={4}
          >
            <Button
              p={0}
              onClick={() => setOnOpenField(!onOpenField)}
              bg={'transparent'}
              border={'none'}
              boxShadow={'none'}
              color={'grayModern.900'}
              fontWeight={400}
              fontSize={'12px'}
              lineHeight={'16px'}
              mr={4}
              leftIcon={
                <MyIcon
                  name="arrowRight"
                  color={'grayModern.500'}
                  w={'16px'}
                  transform={onOpenField ? 'rotate(90deg)' : 'rotate(0)'}
                  transition="transform 0.2s ease"
                />
              }
              _hover={{
                color: 'brightBlue.600',
                '& svg': {
                  color: 'brightBlue.600'
                }
              }}
            >
              {t('logNumber')}
            </Button>
            <Flex alignItems={'center'} gap={2} mr={2}>
              <Text fontSize={'12px'} lineHeight={'16px'} color={'grayModern.500'}>
                {t('visible')}:
              </Text>
              <Text fontSize={'12px'} lineHeight={'16px'} color={'grayModern.500'}>
                {visibleFieldCount} {lang === 'zh' ? t('piece') : ''}
              </Text>
            </Flex>
            <Box h={'12px'} mr={2}>
              <Divider orientation="vertical" color={'grayModern.300'} borderWidth={'1px'} />
            </Box>
            <Flex alignItems={'center'} gap={2}>
              <Text fontSize={'12px'} lineHeight={'16px'} color={'grayModern.500'}>
                {t('hidden')}:
              </Text>
              <Text fontSize={'12px'} lineHeight={'16px'} color={'grayModern.500'}>
                {hiddenFieldCount} {lang === 'zh' ? t('piece') : ''}
              </Text>
            </Flex>
          </Flex>
        </Flex>
        <Button
          minW={'75px'}
          fontSize={'12px'}
          variant={'outline'}
          h={'32px'}
          leftIcon={<MyIcon name="export" />}
        >
          {t('export_log')}
        </Button>
      </Flex>
      {/* fields */}
      <Collapse in={onOpenField} animateOpacity>
        <Flex
          p={4}
          position={'relative'}
          h={'100%'}
          w={'100%'}
          bg={'grayModern.50'}
          borderRadius={'8px'}
        >
          <CheckboxGroup colorScheme="brightBlue">
            {fieldList.map((item) => (
              <Checkbox
                isChecked={item.checked}
                key={item.value}
                onChange={() =>
                  setFieldList(
                    fieldList.map((field) =>
                      field.value === item.value ? { ...field, checked: !field.checked } : field
                    )
                  )
                }
                sx={{
                  'span.chakra-checkbox__control[data-checked]': {
                    background: '#f0f4ff ',
                    border: '1px solid #219bf4 ',
                    boxShadow: '0px 0px 0px 2.4px rgba(33, 155, 244, 0.15)',
                    color: '#219bf4',
                    borderRadius: '4px'
                  }
                }}
              >
                {item.label}
              </Checkbox>
            ))}
          </CheckboxGroup>
        </Flex>
      </Collapse>
    </Flex>
  );
};
