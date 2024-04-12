import { ProductTableProps } from '@/app/components/product-table/price-calculator';
import { FavoriteOptionsProps } from './options';

export interface FavoriteTableProps extends ProductTableProps {
  RetailPrice:number
  storeFee:number
  VAT:number
  cost:number
}

export default class MarginCalculator {
  ProductPrice!: number;

  execute(props: ProductTableProps[], favoratieOptions:FavoriteOptionsProps): FavoriteTableProps[] {
    return props.map((r) => this.calcAll(r, favoratieOptions));
  }
  calcAll(props: ProductTableProps, favoratieOptions:FavoriteOptionsProps): FavoriteTableProps {
    // Y 판매가 X 구매가 a 수수료(%) b 부가세(%) c 마진율(%)
    // 수수료 계산 = aY,부가세 계산 = (Y-X-aY)*b
    this.calcTotalPrice(props);
    const RetailPrice = this.calcRetailPrice(favoratieOptions);
    const storeFee = this.calcStoreFee(RetailPrice, favoratieOptions.storeFee);
    const VAT = this.calcVATFee(RetailPrice - (this.ProductPrice + storeFee), favoratieOptions.VAT);
    const cost = this.ProductPrice + storeFee + VAT;

    return {
      ...props, RetailPrice, storeFee, VAT, cost,
    };
  }

  calcTotalPrice(props:ProductTableProps) {
    const productPrice = props.productPrice.KRWPrice;
    const deliveryPrice = props.delivery.KRWShippingFee;
    const tax = props.tax.totalTax;
    this.ProductPrice = productPrice + deliveryPrice + tax;
  }

  calcRetailPrice(favoratieOptions:FavoriteOptionsProps):number {
    // Y 판매가 X 구매가 a 수수료(%) b 부가세(%) c 마진율(%)
    // 수수료 계산 = aY,부가세 계산 = (Y-X-aY)*b
    // Y = (X + aY +(Y-X-aY)*b)*c
    // Y = ((1-b)/(1/c)-a+(a-1)b)X
    const a = favoratieOptions.storeFee / 100;
    const b = favoratieOptions.VAT / 100;
    const c = 1 + favoratieOptions.margin / 100;

    const numerator = (1 - b);
    const denominator = ((1 / c) - a + (a - 1) * b);
    return (numerator / denominator) * this.ProductPrice;
  }

  calcStoreFee(ProductPrice:number, storeFee:number) {
    return storeFee !== 0 ? ProductPrice * (storeFee / 100) : 0;
  }

  calcVATFee(price:number, VAT:number) {
    return VAT !== 0 ? price * (VAT / 100) : 0;
  }
}
