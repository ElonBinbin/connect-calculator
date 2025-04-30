import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import Calculator from './Calculator';
import { calculatorClient } from '../services/calculatorClient';

// 模拟 calculatorClient
jest.mock('../services/calculatorClient', () => ({
  calculatorClient: {
    add: jest.fn(),
    subtract: jest.fn(),
    multiply: jest.fn(),
    divide: jest.fn(),
  },
}));

describe('计算器组件', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  test('渲染计算器组件', () => {
    render(<Calculator />);
    expect(screen.getByText('使用Go语言调用grpc实现计数器')).toBeInTheDocument();
    expect(screen.getByPlaceholderText('第一个数字')).toBeInTheDocument();
    expect(screen.getByPlaceholderText('第二个数字')).toBeInTheDocument();
    expect(screen.getByText('加法 (+)')).toBeInTheDocument();
    expect(screen.getByText('减法 (-)')).toBeInTheDocument();
    expect(screen.getByText('乘法 (×)')).toBeInTheDocument();
    expect(screen.getByText('除法 (÷)')).toBeInTheDocument();
  });

  test('加法操作', async () => {
    (calculatorClient.add as jest.Mock).mockResolvedValue({ value: 8 });
    render(<Calculator />);
    await userEvent.type(screen.getByPlaceholderText('第一个数字'), '5');
    await userEvent.type(screen.getByPlaceholderText('第二个数字'), '3');
    fireEvent.click(screen.getByText('加法 (+)'));
    await waitFor(() => {
      expect(screen.getByText('计算式: 5 + 3 = 8')).toBeInTheDocument();
      expect(screen.getByText('结果: 8')).toBeInTheDocument();
    });
    expect(calculatorClient.add).toHaveBeenCalledWith({ a: 5, b: 3 });
  });

  test('除法操作 - 除数为零', async () => {
    render(<Calculator />);
    await userEvent.type(screen.getByPlaceholderText('第一个数字'), '10');
    await userEvent.type(screen.getByPlaceholderText('第二个数字'), '0');
    fireEvent.click(screen.getByText('除法 (÷)'));
    expect(screen.getByText('除数不能为零')).toBeInTheDocument();
    expect(calculatorClient.divide).not.toHaveBeenCalled();
  });

  test('输入验证', async () => {
    render(<Calculator />);
    fireEvent.click(screen.getByText('乘法 (×)'));
    expect(screen.getByText('请输入有效的数字')).toBeInTheDocument();
  });
});