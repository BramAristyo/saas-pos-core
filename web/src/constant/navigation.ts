import {
  LayoutDashboard,
  Package,
  Receipt,
  Users,
  Wallet,
} from 'lucide-vue-next'

export const NAVIGATION_CONFIG = [
  {
    title: 'Overview',
    url: '#',
    icon: LayoutDashboard,
    items: [
      {
        title: 'Dashboard',
        url: '/dashboard',
      },
    ],
  },
  {
    title: 'Catalog',
    url: '#',
    icon: Package,
    items: [
      {
        title: 'Categories',
        url: '/categories',
      },
      {
        title: 'Modifiers',
        url: '/modifiers',
      },
    ],
  },
  {
    title: 'Transactions',
    url: '#',
    icon: Receipt,
    items: [
      {
        title: 'Taxes',
        url: '/taxes',
      },
      {
        title: 'Sales Types',
        url: '/sales-types',
      },
      {
        title: 'Discounts',
        url: '/discounts',
      },
    ],
  },
  {
    title: 'Employees',
    url: '#',
    icon: Users,
    items: [
      {
        title: 'List',
        url: '/employees',
      },
      {
        title: 'Shift Schedule',
        url: '/shift-schedules',
      },
      {
        title: 'Attendance',
        url: '/attendances',
      },
      {
        title: 'Payroll',
        url: '/payroll',
      },
    ],
  },
  {
    title: 'Accounting',
    url: '#',
    icon: Wallet,
    items: [
      {
        title: 'Chart of Accounts',
        url: '/coa',
      },
    ],
  },
]
