'use client';

import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import React, { useEffect, useState } from 'react';
import { Column, Header } from '@tanstack/react-table';
import { MagnifyingGlassIcon, CaretSortIcon } from '@radix-ui/react-icons';
import { QuestionToolTip } from '@/components/table-template/utils';
import Select, { components, OptionProps, DropdownIndicatorProps } from 'react-select';
import { Button } from '@/components/ui/button';

export type DefaultSelectProps = {
  value: string,
  label: string,
};

export type SelectFilterProps<V, H> = {
  uniqueValues: V[];
  header: Header<H, any>;
  option: (props: OptionProps<V>) => React.JSX.Element;
};

export interface SelectFilterDialogProps<V, H> extends SelectFilterProps<V, H> {
  columnName:string
}

function DropdownIndicator<T>(props: DropdownIndicatorProps<T>) {
  return (
    <components.DropdownIndicator {...props}>
      <MagnifyingGlassIcon className="h-4 w-4" />
    </components.DropdownIndicator>
  );
}

function closeDialog():void {
  const escapeKeyPressEvent = new KeyboardEvent('keydown', { key: 'Escape' });
  document.dispatchEvent(escapeKeyPressEvent);
}
function AdoptDialog<T>(column: Column<T, any>, selectedValue: any | 'reset') {
  column.setFilterValue(selectedValue !== 'reset' ? selectedValue : undefined);
  closeDialog();
}

export function SelectFilter<V extends DefaultSelectProps, H>(
  { uniqueValues, header, option }:SelectFilterProps<V, H>,
) {
  const [value, setValue] = useState<string[]>([]);

  const currFilterValue:any = header.column.getFilterValue();
  const defaultValue = currFilterValue === undefined
    ? []
    : uniqueValues.filter((v) => currFilterValue.find((f:string) => f === v.value));

  return (
    <div className="flex flex-col h-[420px] justify-between w-full">
      <Select
        onChange={(valArr) => setValue(valArr.map((val) => val.value))}
        options={uniqueValues}
        maxMenuHeight={300}
        defaultValue={defaultValue}
        menuIsOpen
        isMulti
        controlShouldRenderValue={false}
        hideSelectedOptions={false}
        components={{ Option: option, DropdownIndicator }}
        placeholder="검색하기"
        noOptionsMessage={() => '검색 결과가 없습니다.'}
        styles={{
          control: (base) => ({
            ...base,
            '&:hover': { borderColor: 'gray' }, // border style on hover
            border: '1px solid lightgray', // default border color
            boxShadow: 'none', // no box-shadow
            cursor: 'text',
          }),
          option: (base, state) => ({
            backgroundColor: state.isSelected ? '#e2e8f0' : 'white',
          }),
        }}
      />
      <div className="flex gap-2 justify-end">
        <Button asChild={false} variant="outline" onClick={closeDialog}>취소하기</Button>
        <Button asChild={false} onClick={() => AdoptDialog(header.column, value)}>적용하기</Button>
      </div>
    </div>
  );
}

export function SelecFilterDialog<V extends DefaultSelectProps, H>({
  uniqueValues, columnName, header, option,
}
:SelectFilterDialogProps<V, H>) {
  const trigger = (
    <Button asChild={false} variant="ghost">
      {columnName}
      {' '}
      <CaretSortIcon className="w-4 h-4" />
    </Button>
  );
  return (
    <Dialog>
      <DialogTrigger asChild>{trigger}</DialogTrigger>
      <DialogContent className="w-[400px] max-h-[500px]">
        <DialogHeader>
          <DialogTitle>{columnName}</DialogTitle>
          <SelectFilter<V, H>
            uniqueValues={uniqueValues}
            header={header}
            option={option}
          />
        </DialogHeader>
      </DialogContent>
    </Dialog>

  );
}

export type YesOrNoFilterProps<H> = {
  keyString:string
  header: Header<H, any>;
};

function YesOrNoSelect<H>({ keyString, header }:YesOrNoFilterProps<H>) {
  const [selectedValue, setSelectedValue] = useState<{ yes:boolean, no:boolean }>(
    { yes: false, no: false },
  );
  useEffect(() => {
    const filterValue = header.column.getFilterValue();
    let inputState = { yes: false, no: false };
    switch (filterValue) {
      case true:
        inputState = { yes: true, no: false };
        break;
      case false:
        inputState = { yes: false, no: true };
        break;
      default:
        break;
    }

    setSelectedValue(inputState);
  }, [header.column]);

  const yes = `${keyString}-yes`;
  const no = `${keyString}-no`;

  let adoptResult = () => {};

  if (selectedValue.yes && selectedValue.no) {
    adoptResult = () => AdoptDialog(header.column, 'reset');
  }
  if (!selectedValue.yes && !selectedValue.no) {
    adoptResult = () => AdoptDialog(header.column, 'reset');
  }
  if (selectedValue.yes === true && selectedValue.no === false) {
    adoptResult = () => AdoptDialog(header.column, true);
  }
  if (selectedValue.yes === false && selectedValue.no === true) {
    adoptResult = () => AdoptDialog(header.column, false);
  }

  return (
    <>
      <div className="flex justify-between items-center h-[80px] text-base px-2">
        <div>
          <label htmlFor={yes} className="cursor-pointer">
            <input
              id={yes}
              type="checkbox"
              className="accent-black w-5"
              onChange={() => {
                setSelectedValue((obj) => ({ ...obj, yes: !obj.yes }));
              }}
              checked={selectedValue.yes}
            />
            <span className="ml-2">예</span>
          </label>
        </div>
        <div>
          <label htmlFor={no} className="cursor-pointer">
            <input
              id={no}
              type="checkbox"
              className="accent-black w-5"
              onChange={() => {
                setSelectedValue((obj) => ({ ...obj, no: !obj.no }));
              }}
              checked={selectedValue.no}
            />
            <span className="ml-2">아니요</span>
          </label>
        </div>
      </div>
      <div className="flex gap-2 justify-end">
        <Button asChild={false} variant="outline" onClick={closeDialog}>취소하기</Button>
        <Button asChild={false} onClick={adoptResult}>적용하기</Button>
      </div>
    </>
  );
}

export interface YesOrNoFilterDialogProps<H> extends YesOrNoFilterProps<H> {
  hoverCell:string | React.JSX.Element | null
}

export function YesOrNoFilterDialog<H>({
  keyString, header, hoverCell = null,
}:YesOrNoFilterDialogProps<H>) {
  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button asChild={false} variant="ghost">
          {keyString}
          <CaretSortIcon />
          {hoverCell ? <QuestionToolTip hoverCell={hoverCell} /> : null}
        </Button>
      </DialogTrigger>
      <DialogContent className="w-[300px] h-[200px]">
        <DialogHeader>
          <DialogTitle>{keyString}</DialogTitle>
          <YesOrNoSelect keyString={keyString} header={header} />
        </DialogHeader>
      </DialogContent>
    </Dialog>

  );
}
