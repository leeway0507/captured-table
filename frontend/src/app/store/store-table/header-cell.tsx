'use client';

import { Avatar, AvatarImage } from '@/components/ui/avatar';
import { CellContext } from '@tanstack/react-table';
import { IntlShippingFee } from '@/app/type';
import { Button } from '@/components/ui/button';
import Link from 'next/link';
import { CURR, customHoverCard } from '@/components/table-template/utils';
import { KorCurrency, currencySymbol } from '@/app/components/meta/currency';
import { CountryToISO2 } from '../../components/meta/country';
import { StoreTableProps } from './data-preprocessor';

export function Store({ props }: { props: CellContext<StoreTableProps, any> }) {
  const store = props.row.original;
  return (
    <div className="flex justify-between ms-4">
      <Avatar>
        <AvatarImage src={`/store/logo/${store.store_name}.webp`} className="border border-black/40 rounded-full" />
      </Avatar>
      <div className="uppercase flex-col items-center grow">
        <div>{store.store_name.replaceAll('_', ' ')}</div>
        <div className="text-gray-400">
          {store.kor_store_name}
        </div>
      </div>
    </div>
  );
}
export function Country({ props }: { props: CellContext<StoreTableProps, any> }) {
  const store = props.row.original;
  const country = CountryToISO2.find((c) => c.countryCode === store.country);

  return (
    <div className="flex-center gap-2">
      {country!.countryNameKor}
      <div className="text-lg">
        {country!.flag}
      </div>
    </div>

  );
}
export function TaxReduction({ props }: { props: CellContext<StoreTableProps, any> }) {
  const TaxReductionRate:number | undefined = props.row.original.tax_reduction;
  const salePercentFormat = TaxReductionRate?.toLocaleString(
    undefined,
    { style: 'percent', minimumFractionDigits: 0 },
  );

  return TaxReductionRate === undefined ? '0%' : `${salePercentFormat}`;
}

function SingleFeeCell(price:number, currency:string) {
  return CURR(price, currency);
}
function MultipleFeeCell(priceArr:number[], currency:string) {
  const [p1, p2, p3] = priceArr;
  return (
    <div className="grid grid-cols-2">
      <span>의류1</span>
      {CURR(p1, currency)}
      <span>의류2</span>
      {CURR(p2, currency)}
      <span>신발</span>
      {CURR(p3, currency)}
    </div>
  );
}

export function DeliveryFee({ props }: { props: CellContext<StoreTableProps, any> }) {
  const rawShippingFees:IntlShippingFee = props.getValue();
  const currencyCode = props.row.original.currency;
  const shippingFeeArr = Object.values(rawShippingFees);
  const krwShippingFeeArr = Object.values(props.row.original.krw_intl_shipping_fee);

  const singleFee = new Set(shippingFeeArr).size === 1;

  return singleFee
    ? customHoverCard(
      SingleFeeCell(krwShippingFeeArr[0], 'KRW'),
      SingleFeeCell(shippingFeeArr[0], currencyCode),
      'right',
    )
    : customHoverCard(
      MultipleFeeCell(krwShippingFeeArr, 'KRW'),
      MultipleFeeCell(shippingFeeArr, currencyCode),
      'right',
    );
}

export function FreeDeliveryFeeMin({ props }: { props: CellContext<StoreTableProps, any> }) {
  const freeShippingFee:number = props.getValue();
  const currencyCode = props.row.original.currency;
  const KRWFreeshippingFee = props.row.original.krw_intl_free_shipping_min;

  return freeShippingFee !== undefined ? customHoverCard(
    SingleFeeCell(KRWFreeshippingFee, 'KRW'),
    SingleFeeCell(freeShippingFee, currencyCode),
    'right',
  ) : '아니요';
}

export function CurrencyCode({ props }: { props: CellContext<StoreTableProps, any> }) {
  const code:string = props.getValue();
  const korCurrency = KorCurrency[code];
  const symbol = currencySymbol[code];
  return (
    <>
      <div>
        {code}
        (
        {symbol}
        )
      </div>
      <div className="text-xs text-gray-400">{korCurrency}</div>
    </>
  );
}
export function YesOrNo({ props }: { props: CellContext<StoreTableProps, any> }) {
  const isTrue:boolean = props.getValue();
  return isTrue ? '예' : '아니요';
}

export function DeliveryAgency<T>({ props }: { props: CellContext<T, any> }) {
  const s:string = props.getValue();
  const DeliveryAgencyArr = s.split(',');
  const numAgencies = DeliveryAgencyArr.length;
  const cell = (
    <div className="flex-center justify-between" key={DeliveryAgencyArr[0]}>
      <Avatar>
        <AvatarImage
          src={`/delivery_agency/${DeliveryAgencyArr[0].toLowerCase()}.webp`}
          className="border border-black/20 rounded-full"
        />
      </Avatar>
      <div className="uppercase flex-col items-center grow">
        {DeliveryAgencyArr[0]}
        {numAgencies > 1 ? ` + ${numAgencies - 1}` : null}
      </div>
    </div>
  );
  const hoverCell = (
    <>
      {DeliveryAgencyArr.map((agency) => (
        <div className="flex-center justify-between gap-4 mx-2 py-2" key={agency}>
          <Avatar>
            <AvatarImage src={`/delivery_agency/${agency.toLowerCase()}.webp`} className="border border-black/20 rounded-full" />
          </Avatar>
          <div className="uppercase flex-col items-center grow">{agency}</div>
        </div>
      ))}
    </>
  );
  return numAgencies ? customHoverCard(cell, hoverCell) : cell;
}

export function MoveToSite({ props }: { props: CellContext<StoreTableProps, any> }) {
  return (
    <Button variant="secondary" className="font-medium" asChild>
      <Link href={props.getValue()} target="_blank" rel="noreferrer">이동하기</Link>
    </Button>
  );
}
